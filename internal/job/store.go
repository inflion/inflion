package job

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.etcd.io/etcd/clientv3"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Store interface {
	List(ctx context.Context) ([]Job, error)
	Create(ctx context.Context, job Job) error
	Remove(ctx context.Context, job Job) error
}

type EtcdStore struct {
	client *clientv3.Client
}

func (e EtcdStore) List(ctx context.Context) ([]Job, error) {
	projects, err := e.listProjects(ctx)
	if err != nil {
		return nil, err
	}

	log.Println(projects)

	var merged []Job
	for _, p := range projects {
		jobs, err := e.list(ctx, p)
		if err != nil {
			return nil, err
		}
		for _, j := range jobs {
			merged = append(merged, j)
		}
	}
	return merged, nil
}

func (e EtcdStore) Create(ctx context.Context, job Job) error {
	ops := []clientv3.Op{
		clientv3.OpPut(e.scheduleKey(job), job.Schedule),
		clientv3.OpPut(e.flowKey(job), job.FlowId.String()),
	}

	for _, op := range ops {
		_, err := e.connect().Do(ctx, op)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e EtcdStore) Remove(ctx context.Context, job Job) error {
	ops := []clientv3.Op{
		clientv3.OpDelete(e.scheduleKey(job)),
		clientv3.OpDelete(e.flowKey(job)),
	}
	for _, op := range ops {
		_, err := e.connect().Do(ctx, op)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e EtcdStore) listProjects(ctx context.Context) ([]string, error) {
	m := map[string]bool{}
	var projects []string

	resp, err := e.connect().Get(ctx, "/", clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, err
	}

	for _, v := range resp.Kvs {
		m[strings.Split(string(v.Key), "/")[1]] = true
	}
	for k, _ := range m {
		projects = append(projects, k)
	}

	return projects, nil
}

func (e EtcdStore) list(ctx context.Context, project string) ([]Job, error) {
	resp, err := e.connect().Get(ctx, "/"+project+"/jobs", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	jobMap := map[int]Job{}

	for _, kv := range resp.Kvs {
		project := e.extractProjectFrom(kv.Key)
		log.Printf("---- %s ----", string(kv.Key))
		jobId, err := strconv.Atoi(e.extractJobIdFrom(kv.Key))
		if err != nil {
			log.Printf("job error: invalid id format %+v", err)
			continue
		}

		job, ok := jobMap[jobId]

		if !ok {
			job = Job{
				Id:      JobID(jobId),
				Project: project,
			}
		}

		if strings.Contains(string(kv.Key), "/flow") {
			if id, err := uuid.Parse(string(kv.Value)); err == nil {
				job.FlowId = id
			} else {
				log.Printf("job error: invalid uuid format, job id = %d", jobId)
				continue
			}
		} else if strings.Contains(string(kv.Key), "schedule") {
			job.Schedule = string(kv.Value)
		}

		jobMap[jobId] = job
	}

	var jobs []Job
	for _, j := range jobMap {
		jobs = append(jobs, j)
	}

	return jobs, nil
}

func (e EtcdStore) key(job Job) string {
	return fmt.Sprintf("/%s/jobs/%d", job.Project, job.Id)
}

func (e EtcdStore) scheduleKey(job Job) string {
	return fmt.Sprintf("%s/schedule", e.key(job))
}

func (e EtcdStore) flowKey(job Job) string {
	return fmt.Sprintf("%s/flow", e.key(job))
}

func (e EtcdStore) extractJobIdFrom(key []byte) string {
	return strings.Split(string(key), "/")[3]
}

func (e EtcdStore) extractProjectFrom(key []byte) string {
	return strings.Split(string(key), "/")[1]
}

func (e EtcdStore) connect() *clientv3.Client {
	var err error
	if e.client == nil {
		e.client, err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(os.Getenv("ETCD_ENDPOINTS"), ","),
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			log.Println(err)
		}
	}
	return e.client
}

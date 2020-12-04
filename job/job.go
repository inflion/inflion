package job

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/inflion/inflion/flow"
	"github.com/robfig/cron/v3"
	"log"
)

type JobID int

type Job struct {
	Id       JobID
	Project  string
	Schedule string
	FlowId   uuid.UUID
}

type JobRunner interface {
	Run() error
}

type FlowJobRunner struct {
	flowId  uuid.UUID
	project string
	store   flow.Store
}

func NewFlowJobRunner(flowId uuid.UUID, project string, store flow.Store) JobRunner {
	return FlowJobRunner{flowId: flowId, project: project, store: store}
}

func (f FlowJobRunner) Run() error {
	// Run jobs

	return nil
}

type EntryIdJobIdMap map[JobID]cron.EntryID

type JobScheduler interface {
	Start()
	RunOrReplaceJob(job Job) error
	Remove(job Job) error
}

type RealCronScheduler struct {
	cron  *cron.Cron
	idMap EntryIdJobIdMap
}

func NewRealCronScheduler() JobScheduler {
	return RealCronScheduler{cron: cron.New(), idMap: EntryIdJobIdMap{}}
}

func (r RealCronScheduler) Start() {
	r.cron.Start()
}

func (r RealCronScheduler) RunOrReplaceJob(job Job) error {
	entryId, ok := r.idMap[job.Id]
	if ok {
		r.cron.Remove(entryId)
	}

	newEntryId, err := r.cron.AddFunc(job.Schedule, func() {
		runner := NewFlowJobRunner(job.FlowId, job.Project, flow.EtcdBackedFlowStore{})
		err := runner.Run()
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("job craeted %+v", job)

	r.idMap[job.Id] = newEntryId

	return nil
}

func (r RealCronScheduler) Remove(job Job) error {
	entryId, ok := r.idMap[job.Id]
	if ok {
		r.cron.Remove(entryId)
		log.Printf("job removed %d", job.Id)
	} else {
		return fmt.Errorf("job id not found. job id = %d", job.Id)
	}
	return nil
}

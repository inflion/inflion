package store

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type EtcdBackedFlowStore struct {
	client *clientv3.Client
}

func (e EtcdBackedFlowStore) Create(request FlowCreateRequest) (FlowCreateResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return FlowCreateResponse{}, err
	}

	_, err = e.connect().Put(context.Background(), e.createKey(request.Project, id), request.Body)
	if err != nil {
		return FlowCreateResponse{}, err
	}

	return FlowCreateResponse{Id: id}, nil
}

func (e EtcdBackedFlowStore) Get(request FlowGetRequest) (FlowGetResponse, error) {
	key := e.createKey(request.Project, request.Id)
	log.Printf("get by key: %s", key)
	resp, err := e.connect().Get(context.Background(), key)
	if err != nil {
		log.Fatal(err)
		return FlowGetResponse{}, err
	}

	log.Printf("%+v", resp.Kvs)

	return FlowGetResponse{
		Body: string(resp.Kvs[0].Value),
	}, nil
}

func (e EtcdBackedFlowStore) Update(request FlowUpdateRequest) error {
	_, err := e.connect().Put(context.Background(), e.createKey(request.Project, request.Id), request.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e EtcdBackedFlowStore) Delete(request FlowDeleteRequest) error {
	_, err := e.connect().Delete(context.Background(), e.createKey(request.Project, request.Id))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e EtcdBackedFlowStore) createKey(project Project, id FlowId) string {
	return fmt.Sprintf("/%s/flows/%s", project, id.String())
}

func (e EtcdBackedFlowStore) connect() *clientv3.Client {
	var err error
	if e.client == nil {
		e.client, err = clientv3.New(clientv3.Config{
			Endpoints:   []string{"localhost:2379"},
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			log.Println(err)
		}
	}
	return e.client
}

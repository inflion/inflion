package configstore

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"os"
	"strings"
	"time"
)

type EtcdConfigStore struct {
	client *clientv3.Client
}

func (e EtcdConfigStore) CreateOrUpdate(request ConfigCreateOrUpdateRequest) (ConfigCreateOrUpdateResponse, error) {
	_, err := e.connect().Put(context.Background(), e.createKey(request.Project, request.Config.Key), request.Config.Value)
	if err != nil {
		return ConfigCreateOrUpdateResponse{}, err
	}

	return ConfigCreateOrUpdateResponse{Config: Config{
		Key:   request.Config.Key,
		Value: request.Config.Value,
	}}, nil
}

func (e EtcdConfigStore) List(request ConfigListRequest) (ConfigListResponse, error) {
	key := e.createKey(request.Project, request.Key)
	log.Printf("get by key: %s", key)
	resp, err := e.connect().Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
		return ConfigListResponse{}, err
	}

	log.Printf("%+v", resp.Kvs)

	var configs []Config

	for _, v := range resp.Kvs {
		configs = append(configs, Config{Key: string(v.Key), Value: string(v.Value)})
	}

	return ConfigListResponse{
		Configs: configs,
	}, nil
}

func (e EtcdConfigStore) Get(request ConfigGetRequest) (ConfigGetResponse, error) {
	key := e.createKey(request.Project, request.Key)
	log.Printf("get by key: %s", key)
	resp, err := e.connect().Get(context.Background(), key)
	if err != nil {
		log.Fatal(err)
		return ConfigGetResponse{}, err
	}

	log.Printf("%+v", resp.Kvs)

	return ConfigGetResponse{
		Config: Config{
			Key:   key,
			Value: string(resp.Kvs[0].Value),
		},
	}, nil
}

func (e EtcdConfigStore) Delete(request ConfigDeleteRequest) error {
	_, err := e.connect().Delete(context.Background(), e.createKey(request.Project, request.Key))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e EtcdConfigStore) createKey(project Project, key ConfigKey) string {
	return fmt.Sprintf("/%s/config/%s", project, key)
}

func (e EtcdConfigStore) connect() *clientv3.Client {
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

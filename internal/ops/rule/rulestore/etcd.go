// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package rulestore

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/google/uuid"
	"github.com/inflion/inflion/internal/ops/rule"
	"github.com/inflion/inflion/internal/ops/rule/jsonv1"
	"log"
	"os"
	"strings"
	"time"
)

type RuleJson struct {
	Id        uuid.UUID
	ProjectId int64
	Body      json.RawMessage
}

func (r RuleJson) Unmarshal() (rule.Rule, error) {
	j, err := jsonv1.Unmarshal(r.Body)
	if err != nil {
		return rule.Rule{}, err
	}
	return j, nil
}

type EtcdStore struct {
	client *clientv3.Client
}

func (e EtcdStore) createKey(rule RuleJson) string {
	return fmt.Sprintf("/ops/%d/rules/%s", rule.ProjectId, rule.Id.String())
}

func (e EtcdStore) Create(rule RuleJson) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = e.connect().Put(context.Background(), e.createKey(rule), string(rule.Body))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (e EtcdStore) Get(rule RuleJson) (RuleJson, error) {
	resp, err := e.connect().Get(context.Background(), e.createKey(rule))
	if err != nil {
		log.Fatal(err)
	}

	return RuleJson{
		Id:   uuid.UUID{},
		Body: resp.Kvs[0].Value,
	}, nil
}

func (e EtcdStore) Update(rule RuleJson) error {
	_, err := e.connect().Put(context.Background(), e.createKey(rule), string(rule.Body))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e EtcdStore) Delete(rule RuleJson) error {
	_, err := e.connect().Delete(context.Background(), e.createKey(rule))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e EtcdStore) GetRules(projectId int64) ([]rule.Rule, error) {
	key := fmt.Sprintf("/ops/%d/rules", projectId)
	v, err := e.client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rules []rule.Rule

	for _, ruleJson := range v.Kvs {
		r, err := jsonv1.Unmarshal(ruleJson.Value)
		if err != nil {
			log.Println(err)
			continue
		}
		rules = append(rules, r)
	}

	return rules, nil
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

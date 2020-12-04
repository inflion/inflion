// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package matcher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
	"time"
)

type RuleJson struct {
	Id      uuid.UUID
	Project string
	Body    json.RawMessage
}

func (r RuleJson) Unmarshal() (Rule, error) {
	j, err := Unmarshal(r.Body)
	if err != nil {
		return Rule{}, err
	}
	return j, nil
}

func Unmarshal(rawJson []byte) (Rule, error) {
	m := MetadataJson{}
	err := json.Unmarshal(rawJson, &m)
	if err != nil {
		return Rule{}, err
	}

	if m.Metadata.Format.Version == 1 {
		v1, err := UnmarshalV1(rawJson)
		if err != nil {
			return Rule{}, err
		}

		return Rule{
			RuleName:   v1.Body.Name,
			Target:     v1.Body.Target,
			Conditions: Conditions{Conditions: v1.mustConvertConditions()},
		}, nil
	}

	return Rule{}, fmt.Errorf("json version %d not supported", m.Metadata.Format.Version)
}

type EtcdStore struct {
	client *clientv3.Client
}

func (e EtcdStore) createKey(rule RuleJson) string {
	return fmt.Sprintf("/%s/rules/%s", rule.Project, rule.Id.String())
}

func (e EtcdStore) Create(rule RuleJson) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	rule.Id = id

	c := e.connect()
	defer c.Close()

	_, err = c.Put(context.Background(), e.createKey(rule), string(rule.Body))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (e EtcdStore) Get(rule RuleJson) (RuleJson, error) {
	c := e.connect()
	defer c.Close()

	resp, err := c.Get(context.Background(), e.createKey(rule))
	if err != nil {
		log.Fatal(err)
	}

	return RuleJson{
		Id:   uuid.UUID{},
		Body: resp.Kvs[0].Value,
	}, nil
}

func (e EtcdStore) Update(rule RuleJson) error {
	c := e.connect()
	defer c.Close()

	_, err := c.Put(context.Background(), e.createKey(rule), string(rule.Body))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e EtcdStore) Delete(rule RuleJson) error {
	c := e.connect()
	defer c.Close()

	_, err := c.Delete(context.Background(), e.createKey(rule))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e EtcdStore) GetRules(project string) ([]Rule, error) {
	key := fmt.Sprintf("/%s/rules", project)

	c := e.connect()
	defer c.Close()

	v, err := c.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rules []Rule

	for _, ruleJson := range v.Kvs {
		r, err := Unmarshal(ruleJson.Value)
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
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(os.Getenv("ETCD_ENDPOINTS"), ","),
		DialTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Println(err)
	}
	return e.client
}

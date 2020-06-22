// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package clientv1

import (
	"context"
	pb "github.com/inflion/inflion/inflionserver/rule/rulepb"
	"log"
)

type Rule struct {
	Id   string
	Body string
}

type RuleClient interface {
	Create(body string) (Rule, error)
	Get(id string) (Rule, error)
	Update(id string, body string) (string, error)
	Remove(id string) (string, error)
}

type RuleClientPb struct {
	project  string
	endpoint string
}

func NewRuleClient(project string, endpoint string) RuleClient {
	return RuleClientPb{project: project, endpoint: endpoint}
}

func (f RuleClientPb) Create(body string) (Rule, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		log.Println(err)
		return Rule{}, err
	}

	client := pb.NewRuleClient(c.conn)
	res, err := client.Create(context.Background(), &pb.CreateRuleRequest{
		Project: f.project,
		Body:    body,
	})
	if err != nil {
		log.Print(err)
		return Rule{}, nil
	}

	return Rule{
		Id:   res.Id,
		Body: body,
	}, nil
}

func (f RuleClientPb) Get(id string) (Rule, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		log.Println(err)
	}

	client := pb.NewRuleClient(c.conn)
	res, err := client.Get(context.Background(), &pb.GetRuleRequest{
		Project: f.project,
		Id:      id,
	})
	if err != nil {
		log.Print(err)
		return Rule{Id: id}, err
	}

	return Rule{
		Id:   id,
		Body: res.Body,
	}, nil
}

func (f RuleClientPb) Update(id string, body string) (string, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		log.Println(err)
	}

	client := pb.NewRuleClient(c.conn)
	_, err = client.Update(context.Background(), &pb.UpdateRuleRequest{
		Project: f.project,
		Id:      id,
		Body:    body,
	})
	if err != nil {
		return id, err
	}
	return id, nil
}

func (f RuleClientPb) Remove(id string) (string, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		log.Println(err)
		return id, err
	}

	client := pb.NewRuleClient(c.conn)
	res, err := client.Delete(context.Background(), &pb.DeleteRuleRequest{
		Project: f.project,
		Id:      id,
	})
	if err != nil {
		log.Print(err)
		return id, err
	}

	return res.Id, nil
}

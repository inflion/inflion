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
	"fmt"
	pb "github.com/inflion/inflion/inflionserver/inflionserverpb"
	"log"
)

type Flow struct {
	Id   string
	Body string
}

type Flowclient interface {
	Run(id string) (string, error)
	Create(body string) (Flow, error)
	Get(id string) (Flow, error)
	Update(id string, body string) (string, error)
	Remove(id string) (string, error)
}

type Flowclientpb struct {
	project  string
	endpoint string
}

func NewFlowClient(project string, endpoint string) Flowclient {
	return Flowclientpb{project: project, endpoint: endpoint}
}

func (f Flowclientpb) Run(id string) (string, error) {
	c := grpcConnection{
		endpoint: f.endpoint,
	}
	err := c.connect()
	if err != nil {
		return "", err
	}

	client := pb.NewFlowClient(c.conn)
	res, err := client.Run(context.Background(), &pb.RunFlowRequest{
		Project: f.project,
		Id:      id,
	})
	if err != nil {
		return "", err
	}

	if res.Status == pb.RunFlowResponse_FAILURE {
		return "", fmt.Errorf(res.Output)
	}

	return res.Output, nil
}

func (f Flowclientpb) Create(body string) (Flow, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		log.Println(err)
		return Flow{}, err
	}

	client := pb.NewFlowClient(c.conn)
	res, err := client.Create(context.Background(), &pb.CreateFlowRequest{
		Project: f.project,
		Body:    body,
	})
	if err != nil {
		log.Print(err)
		return Flow{}, nil
	}

	return Flow{
		Id:   res.Id,
		Body: body,
	}, nil
}

func (f Flowclientpb) Get(id string) (Flow, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		log.Println(err)
	}

	client := pb.NewFlowClient(c.conn)
	res, err := client.Get(context.Background(), &pb.GetFlowRequest{
		Project: f.project,
		Id:      id,
	})
	if err != nil {
		log.Print(err)
		return Flow{Id: id}, err
	}

	return Flow{
		Id:   id,
		Body: res.Body,
	}, nil
}

func (f Flowclientpb) Update(id string, body string) (string, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		log.Println(err)
	}

	client := pb.NewFlowClient(c.conn)
	_, err = client.Update(context.Background(), &pb.UpdateFlowRequest{
		Project: f.project,
		Id:      id,
		Body:    body,
	})
	if err != nil {
		return id, err
	}
	return id, nil
}

func (f Flowclientpb) Remove(id string) (string, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		log.Println(err)
		return id, err
	}

	client := pb.NewFlowClient(c.conn)
	res, err := client.Delete(context.Background(), &pb.DeleteFlowRequest{
		Project: f.project,
		Id:      id,
	})
	if err != nil {
		log.Print(err)
		return id, err
	}

	return res.Id, nil
}

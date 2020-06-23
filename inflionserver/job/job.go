// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package job

import (
	"context"
	spb "github.com/inflion/inflion/inflionserver/job/jobpb"
	cpb "github.com/inflion/inflion/jobserver/jobserverpb"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

type JobServer struct {
	endpoint string
}

func NewJobServer() JobServer {
	return JobServer{
		endpoint: os.Getenv("JOB_SERVER_ENDPOINT"),
	}
}

type grpcConnection struct {
	conn     *grpc.ClientConn
	endpoint string
}

func (c *grpcConnection) connect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	c.conn, err = grpc.DialContext(ctx, c.endpoint, grpc.WithInsecure(), grpc.WithBlock())
	return err
}

func (c *grpcConnection) close() error {
	return c.conn.Close()
}

func (j JobServer) List(ctx context.Context, request *spb.ListJobsRequest) (*spb.ListJobsResponse, error) {
	c := grpcConnection{
		endpoint: j.endpoint,
	}
	err := c.connect()
	if err != nil {
		return nil, err
	}

	jc := cpb.NewJobStoreClient(c.conn)

	r, err := jc.List(ctx, &cpb.ListJobsRequest{
		Project: request.Project,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := &spb.ListJobsResponse{}

	for _, j := range r.Jobs {
		resp.Jobs = append(resp.Jobs, &spb.Job{
			Id:       j.Id,
			Project:  j.Project,
			FlowId:   j.FlowId,
			Schedule: j.Schedule,
		})
	}

	return resp, nil
}

func (j JobServer) Create(ctx context.Context, request *spb.CreateJobRequest) (*spb.CreateJobResponse, error) {
	log.Println("job create")
	c := grpcConnection{
		endpoint: j.endpoint,
	}
	err := c.connect()
	if err != nil {
		return nil, err
	}

	jc := cpb.NewJobStoreClient(c.conn)

	_, err = jc.Create(ctx, &cpb.CreateJobRequest{
		Id:       request.Id,
		Project:  request.Project,
		FlowId:   request.FlowId,
		Schedule: request.Schedule,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &spb.CreateJobResponse{
		Id: "created",
	}, nil
}

func (j JobServer) Remove(ctx context.Context, request *spb.RemoveJobRequest) (*spb.RemoveJobResponse, error) {
	log.Println("remove job")
	c := grpcConnection{
		endpoint: j.endpoint,
	}
	err := c.connect()
	if err != nil {
		return nil, err
	}

	jc := cpb.NewJobStoreClient(c.conn)

	r, err := jc.Remove(ctx, &cpb.RemoveJobRequest{
		Id:      request.Id,
		Project: request.Project,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &spb.RemoveJobResponse{
		Id: r.Id,
	}, nil
}

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
	"github.com/google/uuid"
	pb "github.com/inflion/inflion/inflionserver/job/jobpb"
	"github.com/inflion/inflion/job"
)

type Job struct {
	Id       int32
	Project  string
	FlowId   string
	Schedule string
}

type JobClient interface {
	List(project string) ([]job.Job, error)
	Create(job Job) error
	Remove(job Job) error
}

type JobClientPb struct {
	endpoint string
}

func NewJobClient(endpoint string) JobClient {
	return JobClientPb{endpoint: endpoint}
}

func (f JobClientPb) List(project string) ([]job.Job, error) {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		return nil, err
	}

	client := pb.NewJobInfoClient(c.conn)
	r, err := client.List(context.Background(), &pb.ListJobsRequest{
		Project: project,
	})
	if err != nil {
		return nil, err
	}

	var jobs []job.Job

	for _, rj := range r.Jobs {
		jobs = append(jobs, job.Job{
			Id:       job.JobID(rj.Id),
			Project:  rj.Project,
			Schedule: rj.Schedule,
			FlowId:   uuid.MustParse(rj.FlowId),
		})
	}

	return jobs, nil
}

func (f JobClientPb) Create(job Job) error {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		return err
	}

	client := pb.NewJobInfoClient(c.conn)
	_, err = client.Create(context.Background(), &pb.CreateJobRequest{
		Id:       job.Id,
		Project:  job.Project,
		FlowId:   job.FlowId,
		Schedule: job.Schedule,
	})
	if err != nil {
		return err
	}

	return nil
}

func (f JobClientPb) Remove(job Job) error {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		return err
	}

	client := pb.NewJobInfoClient(c.conn)
	_, err = client.Remove(context.Background(), &pb.RemoveJobRequest{
		Id:      job.Id,
		Project: job.Project,
	})
	if err != nil {
		return err
	}

	return nil
}

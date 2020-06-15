package clientv1

import (
	"context"
	pb "github.com/inflion/inflion/inflionserver/inflionserverpb"
)

type Job struct {
	Id       int32
	Project  string
	FlowId   string
	Schedule string
}

type JobClient interface {
	Create(job Job) error
	Remove(job Job) error
}

type JobClientPb struct {
	endpoint string
}

func NewJobClient(endpoint string) JobClient {
	return JobClientPb{endpoint: endpoint}
}

func (f JobClientPb) Create(job Job) error {
	c := grpcConnection{endpoint: f.endpoint}
	err := c.connect()
	if err != nil {
		return err
	}

	client := pb.NewJobClient(c.conn)
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

	client := pb.NewJobClient(c.conn)
	_, err = client.Remove(context.Background(), &pb.RemoveJobRequest{
		Id:      job.Id,
		Project: job.Project,
	})
	if err != nil {
		return err
	}

	return nil
}

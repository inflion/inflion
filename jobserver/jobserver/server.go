package jobserver

import (
	"context"
	"github.com/google/uuid"
	"github.com/inflion/inflion/internal/job"
	pb "github.com/inflion/inflion/jobserver/jobserverpb"
	"log"
)

type JobServer struct {
	store     job.Store
	scheduler job.JobScheduler
}

func NewJobServer(store job.Store, scheduler job.JobScheduler) JobServer {
	return JobServer{
		store:     store,
		scheduler: scheduler,
	}
}

func (j JobServer) Create(ctx context.Context, request *pb.CreateJobRequest) (*pb.CreateJobResponse, error) {
	log.Println("job create request")

	flowId, err := uuid.Parse(request.FlowId)
	if err != nil {
		return nil, err
	}

	job := job.Job{
		Id:       job.JobID(request.Id),
		Project:  request.Project,
		Schedule: request.Schedule,
		FlowId:   flowId,
	}

	err = j.store.Create(ctx, job)
	if err != nil {
		return nil, err
	}

	err = j.scheduler.RunOrReplaceJob(job)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateJobResponse{
		Id: "created",
	}, nil
}

func (j JobServer) Remove(ctx context.Context, request *pb.RemoveJobRequest) (*pb.RemoveJobResponse, error) {
	job := job.Job{
		Id:      job.JobID(request.Id),
		Project: request.Project,
	}
	err := j.scheduler.Remove(job)
	if err != nil {
		log.Printf("failed to remove a job from scheduler: %+v", err)
		return nil, err
	}

	err = j.store.Remove(ctx, job)
	if err != nil {
		log.Printf("failed to remove a job from store: %+v", err)
		return nil, err
	}

	return &pb.RemoveJobResponse{
		Id: int32(job.Id),
	}, nil
}

//go:generate protoc -I jobserverpb/ jobserverpb/job.proto --go_out=plugins=grpc:jobserverpb
package main

import (
	"context"
	"github.com/inflion/inflion/internal/job"
	"github.com/inflion/inflion/jobserver/jobserver"
	pb "github.com/inflion/inflion/jobserver/jobserverpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const bind = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatal(err)
	}

	scheduler := job.NewRealCronScheduler()
	store := job.EtcdStore{}

	jobs, err := store.List(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, j := range jobs {
		err := scheduler.RunOrReplaceJob(j)
		if err != nil {
			log.Println(err)
		}
	}

	s := grpc.NewServer()
	js := jobserver.NewJobServer(
		job.EtcdStore{},
		scheduler,
	)

	pb.RegisterJobStoreServer(s, js)

	scheduler.Start()

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

//go:generate protoc -I jobserverpb/ jobserverpb/jobserver.proto --go_out=plugins=grpc:jobserverpb
// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

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

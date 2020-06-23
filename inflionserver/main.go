//go:generate protoc -I ../proto/inflion/inflionserver/flow/v1  ../proto/inflion/inflionserver/flow/v1/flow.proto   -I ../proto --go_out=plugins=grpc:flow
//go:generate protoc -I ../proto/inflion/inflionserver/rule/v1  ../proto/inflion/inflionserver/rule/v1/rule.proto   -I ../proto --go_out=plugins=grpc:rule
//go:generate protoc -I ../proto/inflion/inflionserver/job/v1   ../proto/inflion/inflionserver/job/v1/job.proto     -I ../proto --go_out=plugins=grpc:job
//go:generate protoc -I ../proto/inflion/inflionserver/event/v1 ../proto/inflion/inflionserver/event/v1/event.proto -I ../proto --go_out=plugins=grpc:event
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
	"github.com/inflion/inflion/inflionserver/event"
	"github.com/inflion/inflion/inflionserver/event/eventpb"
	"github.com/inflion/inflion/inflionserver/flow"
	"github.com/inflion/inflion/inflionserver/flow/flowpb"
	"github.com/inflion/inflion/inflionserver/job"
	"github.com/inflion/inflion/inflionserver/job/jobpb"
	ruleserver "github.com/inflion/inflion/inflionserver/rule"
	"github.com/inflion/inflion/inflionserver/rule/rulepb"
	"github.com/inflion/inflion/internal/ops/flow/store"
	"github.com/inflion/inflion/internal/ops/rule"
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

	server := grpc.NewServer()

	flowpb.RegisterFlowInfoServer(server, flow.DefaultFlowServer{Store: store.EtcdBackedFlowStore{}})
	rulepb.RegisterRuleServer(server, ruleserver.DefaultRuleServer{Store: rule.EtcdStore{}})
	jobpb.RegisterJobInfoServer(server, job.NewJobServer())
	eventpb.RegisterEventServer(server, event.DefaultEventServer{})

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

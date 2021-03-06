//go:generate protoc -I ../proto/inflion/inflionserver/flow/v1  ../proto/inflion/inflionserver/flow/v1/flow.proto   -I ../proto --go_out=plugins=grpc:flow
//go:generate protoc -I ../proto/inflion/inflionserver/rule/v1  ../proto/inflion/inflionserver/rule/v1/rule.proto   -I ../proto --go_out=plugins=grpc:rule
//go:generate protoc -I ../proto/inflion/inflionserver/job/v1   ../proto/inflion/inflionserver/job/v1/job.proto     -I ../proto --go_out=plugins=grpc:job
//go:generate protoc -I ../proto/inflion/inflionserver/event/v1 ../proto/inflion/inflionserver/event/v1/event.proto -I ../proto --go_out=plugins=grpc:event
//go:generate protoc -I ../proto/inflion/inflionserver/byteevent/v1 ../proto/inflion/inflionserver/byteevent/v1/byteevent.proto -I ../proto --go_out=plugins=grpc:byteevent
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package main

import (
	"github.com/inflion/inflion/flow"
	"github.com/inflion/inflion/flow/matcher"
	"github.com/inflion/inflion/inflionserver/byteevent"
	"github.com/inflion/inflion/inflionserver/byteevent/byteeventpb"
	flowserver "github.com/inflion/inflion/inflionserver/flow"
	"github.com/inflion/inflion/inflionserver/flow/flowpb"
	"github.com/inflion/inflion/inflionserver/job"
	"github.com/inflion/inflion/inflionserver/job/jobpb"
	ruleserver "github.com/inflion/inflion/inflionserver/rule"
	"github.com/inflion/inflion/inflionserver/rule/rulepb"
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

	flowpb.RegisterFlowInfoServer(server, flowserver.DefaultFlowServer{Store: flow.EtcdBackedFlowStore{}})
	rulepb.RegisterRuleServer(server, ruleserver.DefaultRuleServer{Store: matcher.EtcdStore{}})
	jobpb.RegisterJobInfoServer(server, job.NewJobServer())
	byteeventpb.RegisterByteEventServer(server, byteevent.DefaultByteEventServer{})

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

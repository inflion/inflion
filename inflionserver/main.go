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
	"github.com/inflion/inflion/inflionserver/flow"
	"github.com/inflion/inflion/inflionserver/job"
	"github.com/inflion/inflion/inflionserver/rule"
	"github.com/inflion/inflion/internal/ops/flow/store"
	"github.com/inflion/inflion/internal/ops/rule/rulestore"
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

	s := grpc.NewServer()

	flow.RegisterFlowServer(s, flow.DefaultFlowServer{Store: store.EtcdBackedFlowStore{}})
	rule.RegisterRuleServer(s, rule.DefaultRuleServer{Store: rulestore.EtcdStore{}})
	job.RegisterJobInfoServer(s, job.NewJobServer())
	event.RegisterEventServer(s, event.NewDefaultEventServer())

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

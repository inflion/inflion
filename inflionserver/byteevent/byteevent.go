// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package byteevent

import (
	"context"
	"github.com/inflion/inflion/flow"
	"github.com/inflion/inflion/flow/event"
	"github.com/inflion/inflion/flow/matcher"
	pb "github.com/inflion/inflion/inflionserver/byteevent/byteeventpb"
	"log"
)

type DefaultByteEventServer struct{}

func (f DefaultByteEventServer) Put(_ context.Context, request *pb.PutByteEventRequest) (*pb.PutEventResponse, error) {
	e, err := event.NewInflionEvent(request.Project, request.Event)
	if err != nil {
		return nil, err
	}

	p := flow.NewEventProcessor(flow.NewFlowStore(), matcher.NewEventMatcher(matcher.NewRuleStore()))
	if err = p.Process(*e); err != nil {
		log.Println(err)
		return &pb.PutEventResponse{}, err
	}

	return &pb.PutEventResponse{}, nil
}

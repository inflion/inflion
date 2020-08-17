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
	pb "github.com/inflion/inflion/inflionserver/byteevent/byteeventpb"
	inflionEvent "github.com/inflion/inflion/internal/ops/event"
	"github.com/inflion/inflion/internal/ops/producer"
	"log"
)

type DefaultByteEventServer struct{}

func (f DefaultByteEventServer) Put(_ context.Context, request *pb.PutByteEventRequest) (*pb.PutEventResponse, error) {
	log.Print("event request received")
	e, err := inflionEvent.NewInflionEvent(request.Project, request.Event)
	if err != nil {
		return nil, err
	}
	return &pb.PutEventResponse{}, producer.NewProducer().Produce(*e)
}

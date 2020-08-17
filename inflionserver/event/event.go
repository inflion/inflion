// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package event

import (
	"context"
	pb "github.com/inflion/inflion/inflionserver/event/eventpb"
	inflionEvent "github.com/inflion/inflion/internal/ops/event"
	"github.com/inflion/inflion/internal/ops/producer"
)

type DefaultEventServer struct{}

func (f DefaultEventServer) Put(_ context.Context, request *pb.PutEventRequest) (*pb.PutEventResponse, error) {
	ie, err := inflionEvent.NewInflionEvent(request.Project, request.Event)
	if err != nil {
		return nil, err
	}

	return &pb.PutEventResponse{}, producer.NewProducer().Produce(*ie)
}

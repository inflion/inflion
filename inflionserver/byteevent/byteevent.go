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
	"github.com/inflion/inflion/internal/ops/producer"
)

type DefaultByteEventServer struct {
}

func (f DefaultByteEventServer) Put(_ context.Context, request *pb.PutByteEventRequest) (*pb.PutEventResponse, error) {
	return &pb.PutEventResponse{}, producer.NewProducer().Produce(request.Event)
}

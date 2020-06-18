//go:generate protoc -I ../../proto/inflion/inflionserver/event/v1 ../../proto/inflion/inflionserver/event/v1/event.proto -I ../../proto --go_out=plugins=grpc:$GOPATH/src
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
	"fmt"
	"github.com/golang/protobuf/ptypes"
)

type DefaultEventServer struct {
}

func NewDefaultEventServer() *DefaultEventServer {
	return &DefaultEventServer{}
}

func (e DefaultEventServer) Push(ctx context.Context, event *CloudWatchEvent) (*EventResponse, error) {
	details := `{
		"a":12345,
		"b":12345.678,
		"c":"12345",
		"d":true,
		"e":null
		}`

	mockEvent := CloudWatchEvent{
		Version:    "",
		Id:         "",
		DetailType: "",
		Source:     "",
		AccountId:  "",
		Time:       ptypes.TimestampNow(),
		Region:     "",
		Resources: []string{
			"",
		},
		Detail: []byte(details),
	}

	fmt.Print(mockEvent)

	// TODO execute flow associated with event

	return &EventResponse{}, nil
}

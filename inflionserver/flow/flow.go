//go:generate protoc -I ../../proto/inflion/inflionserver/flow/v1 ../../proto/inflion/inflionserver/flow/v1/flow.proto -I ../../proto --go_out=plugins=grpc:$GOPATH/src
// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package flow

import (
	"context"
	"github.com/google/uuid"
	"github.com/inflion/inflion/internal/ops/flow"
	"github.com/inflion/inflion/internal/ops/flow/store"
	"log"
)

type DefaultFlowServer struct {
	Store store.Store
}

func (f DefaultFlowServer) Run(_ context.Context, request *RunFlowRequest) (*RunFlowResponse, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	ec := flow.NewExecutionContext()
	ec.AddFields("system", flow.ExecutionFields{
		Values: map[string]interface{}{
			"project": request.Project,
		},
	})

	opsflow := flow.NewOpsFlow(store.NewStoreRecipeReader(request.Project, id, f.Store))
	result, err := opsflow.Run(ec)
	if err != nil {
		return nil, err
	}

	return &RunFlowResponse{
		Status: RunFlowResponse_SUCCESS,
		Output: result.Message,
	}, nil
}

func (f DefaultFlowServer) Create(_ context.Context, request *CreateFlowRequest) (*CreateFlowResponse, error) {
	r, err := f.Store.Create(store.FlowCreateRequest{
		Project: request.Project,
		Body:    request.Body,
	})
	if err != nil {
		return nil, err
	}

	return &CreateFlowResponse{
		Id: r.Id.String(),
	}, nil
}

func (f DefaultFlowServer) Get(_ context.Context, request *GetFlowRequest) (*GetFlowResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	r, err := f.Store.Get(store.FlowGetRequest{
		Project: request.Project,
		Id:      id,
	})
	if err != nil {
		return nil, err
	}

	return &GetFlowResponse{
		Id:   request.Id,
		Body: r.Body,
	}, nil
}

func (f DefaultFlowServer) Update(_ context.Context, request *UpdateFlowRequest) (*UpdateFlowResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	err = f.Store.Update(store.FlowUpdateRequest{
		Project: request.Project,
		Id:      id,
		Body:    request.Body,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &UpdateFlowResponse{
		Id: id.String(),
	}, nil
}

func (f DefaultFlowServer) Delete(_ context.Context, request *DeleteFlowRequest) (*DeleteFlowResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	err = f.Store.Delete(store.FlowDeleteRequest{
		Project: request.Project,
		Id:      id,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &DeleteFlowResponse{
		Id: id.String(),
	}, nil
}

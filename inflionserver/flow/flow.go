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
	pb "github.com/inflion/inflion/inflionserver/inflionserverpb"
	"github.com/inflion/inflion/internal/ops/flow"
	"github.com/inflion/inflion/internal/ops/flow/store"
	"log"
)

type DefaultFlowServer struct {
	Store store.Store
}

func (f DefaultFlowServer) Run(_ context.Context, request *pb.RunFlowRequest) (*pb.RunFlowResponse, error) {
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

	return &pb.RunFlowResponse{
		Status: pb.RunFlowResponse_SUCCESS,
		Output: result.Message,
	}, nil
}

func (f DefaultFlowServer) Create(_ context.Context, request *pb.CreateFlowRequest) (*pb.CreateFlowResponse, error) {
	r, err := f.Store.Create(store.FlowCreateRequest{
		Project: request.Project,
		Body:    request.Body,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateFlowResponse{
		Id: r.Id.String(),
	}, nil
}

func (f DefaultFlowServer) Get(_ context.Context, request *pb.GetFlowRequest) (*pb.GetFlowResponse, error) {
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

	return &pb.GetFlowResponse{
		Id:   request.Id,
		Body: r.Body,
	}, nil
}

func (f DefaultFlowServer) Update(_ context.Context, request *pb.UpdateFlowRequest) (*pb.UpdateFlowResponse, error) {
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

	return &pb.UpdateFlowResponse{
		Id: id.String(),
	}, nil
}

func (f DefaultFlowServer) Delete(_ context.Context, request *pb.DeleteFlowRequest) (*pb.DeleteFlowResponse, error) {
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

	return &pb.DeleteFlowResponse{
		Id: id.String(),
	}, nil
}

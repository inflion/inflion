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
	pb "github.com/inflion/inflion/inflionserver/flow/flowpb"
	"github.com/inflion/inflion/internal/ops/flow"
	flowContext "github.com/inflion/inflion/internal/ops/flow/context"
	"github.com/inflion/inflion/internal/ops/flow/store"
	"log"
)

type DefaultFlowServer struct {
	Store store.Store
}

func (s DefaultFlowServer) List(_ context.Context, request *pb.ListFlowRequest) (*pb.ListFlowResponse, error) {
	result := &pb.ListFlowResponse{
		Flows: []*pb.Flow{},
	}

	flows, err := s.Store.List(request.Project)
	if err != nil {
		return nil, err
	}

	for _, f := range flows {
		result.Flows = append(result.Flows, &pb.Flow{
			Project: f.Project,
			Id:      f.Id.String(),
			Body:    f.Body,
		})
	}

	return result, nil
}

func (s DefaultFlowServer) Run(_ context.Context, request *pb.RunFlowRequest) (*pb.RunFlowResponse, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}

	ec := flowContext.NewExecutionContext()
	ec.AddField("system", map[string]interface{}{"project": request.Project})

	storedFlow, err := store.NewStoreRecipeReader(request.Project, id, s.Store).Read()
	if err != nil {
		return nil, err
	}
	opsflow := flow.NewFlowExecutor(storedFlow, flow.NewAggregateActionLoader())
	result, err := opsflow.Run(ec)
	if err != nil {
		return nil, err
	}

	return &pb.RunFlowResponse{
		Status: pb.RunFlowResponse_SUCCESS,
		Output: result.Message,
	}, nil
}

func (s DefaultFlowServer) Create(_ context.Context, request *pb.CreateFlowRequest) (*pb.CreateFlowResponse, error) {
	r, err := s.Store.Create(store.FlowCreateRequest{
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

func (s DefaultFlowServer) Get(_ context.Context, request *pb.GetFlowRequest) (*pb.GetFlowResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	r, err := s.Store.Get(store.FlowGetRequest{
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

func (s DefaultFlowServer) Update(_ context.Context, request *pb.UpdateFlowRequest) (*pb.UpdateFlowResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	err = s.Store.Update(store.FlowUpdateRequest{
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

func (s DefaultFlowServer) Delete(_ context.Context, request *pb.DeleteFlowRequest) (*pb.DeleteFlowResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	err = s.Store.Delete(store.FlowDeleteRequest{
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

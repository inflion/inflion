// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package rule

import (
	"context"
	"github.com/google/uuid"
	rule "github.com/inflion/inflion/flow/matcher"
	pb "github.com/inflion/inflion/inflionserver/rule/rulepb"
	"log"
)

type DefaultRuleServer struct {
	Store rule.Store
}

func (f DefaultRuleServer) Create(ctx context.Context, request *pb.CreateRuleRequest) (*pb.CreateRuleResponse, error) {
	id, err := f.Store.Create(rule.RuleJson{
		Project: request.Project,
		Body:    []byte(request.Body),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateRuleResponse{
		Id: id.String(),
	}, nil
}

func (f DefaultRuleServer) Get(ctx context.Context, request *pb.GetRuleRequest) (*pb.GetRuleResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	ruleJson, err := f.Store.Get(rule.RuleJson{
		Project: request.Project,
		Id:      id,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetRuleResponse{
		Id:   request.Id,
		Body: string(ruleJson.Body),
	}, nil
}

func (f DefaultRuleServer) Update(ctx context.Context, request *pb.UpdateRuleRequest) (*pb.UpdateRuleResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	err = f.Store.Update(rule.RuleJson{
		Id:      id,
		Project: request.Project,
		Body:    []byte(request.Body),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.UpdateRuleResponse{
		Id: id.String(),
	}, nil
}

func (f DefaultRuleServer) Delete(ctx context.Context, request *pb.DeleteRuleRequest) (*pb.DeleteRuleResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	err = f.Store.Delete(rule.RuleJson{
		Id: id,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.DeleteRuleResponse{
		Id: id.String(),
	}, nil
}

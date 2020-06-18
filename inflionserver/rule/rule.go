//go:generate protoc -I ../../proto/inflion/inflionserver/rule/v1 ../../proto/inflion/inflionserver/rule/v1/rule.proto -I ../../proto --go_out=plugins=grpc:$GOPATH/src
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
	"github.com/inflion/inflion/internal/ops/rule/rulestore"
	"log"
)

type DefaultRuleServer struct {
	Store rulestore.Store
}

func (f DefaultRuleServer) Create(ctx context.Context, request *CreateRuleRequest) (*CreateRuleResponse, error) {
	id, err := f.Store.Create(rulestore.RuleJson{
		Body: []byte(request.Body),
	})
	if err != nil {
		return nil, err
	}

	return &CreateRuleResponse{
		Id: id.String(),
	}, nil
}

func (f DefaultRuleServer) Get(ctx context.Context, request *GetRuleRequest) (*GetRuleResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	ruleJson, err := f.Store.Get(rulestore.RuleJson{
		Id:        id,
		ProjectId: 0, // TODO project id
	})
	if err != nil {
		return nil, err
	}

	return &GetRuleResponse{
		Id:   request.Id,
		Body: string(ruleJson.Body),
	}, nil
}

func (f DefaultRuleServer) Update(ctx context.Context, request *UpdateRuleRequest) (*UpdateRuleResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	err = f.Store.Update(rulestore.RuleJson{
		Id:   id,
		Body: []byte(request.Body),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &UpdateRuleResponse{
		Id: id.String(),
	}, nil
}

func (f DefaultRuleServer) Delete(ctx context.Context, request *DeleteRuleRequest) (*DeleteRuleResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	err = f.Store.Delete(rulestore.RuleJson{
		Id: id,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &DeleteRuleResponse{
		Id: id.String(),
	}, nil
}

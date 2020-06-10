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
	"encoding/json"
)

func UnmarshalV1(v1FormattedJson []byte) (*FlowJsonV1, error) {
	v1 := FlowJsonV1{}
	err := json.Unmarshal(v1FormattedJson, &v1)
	if err != nil {
		return nil, err
	}
	return &v1, nil
}

type FlowJsonV1 struct {
	Metadata MetadataJson `json:"metadata"`
	Body     struct {
		Conditions []ConditionJsonV1 `json:"conditions"`
		Stages     []StageJsonV1     `json:"stages"`
	} `json:"body"`
}

type ExpressionJsonV1 struct {
	Input     string `json:"input"`
	Operation string `json:"operation"`
	Value     string `json:"value"`
}

func (e *ExpressionJsonV1) mustConvert() Expression {
	return Expression{
		Input:     e.Input,
		Operation: e.Operation,
		Value:     e.Value,
	}
}

type ConditionJsonV1 struct {
	Id          string `json:"id"`
	Expressions []ExpressionJsonV1
	IfTrue      struct {
		NextId string `json:"next_id"`
	} `json:"if_true"`
	IfFalse struct {
		NextId string `json:"next_id"`
	} `json:"if_false"`
}

func (c *ConditionJsonV1) mustConvert() Condition {
	return Condition{
		Id:          c.Id,
		Expressions: c.mustConvertExpressions(),
		IfTrue: NextStage{
			Id:   c.IfTrue.NextId,
			Node: nil,
		},
		IfFalse: NextStage{
			Id:   c.IfFalse.NextId,
			Node: nil,
		},
	}
}

func (c *ConditionJsonV1) mustConvertExpressions() []Expression {
	var e []Expression
	for _, ex := range c.Expressions {
		e = append(e, ex.mustConvert())
	}
	return e
}

type StageJsonV1 struct {
	Id      string              `json:"id"`
	Next    string              `json:"next"`
	Name    string              `json:"name"`
	Actions []StageActionJsonV1 `json:"actions"`
}

type StageActionJsonV1 struct {
	Type   string            `json:"type"`
	Params map[string]string `json:"params"`
}

func (v StageActionJsonV1) mustConvert() Action {
	return Action{
		Type:   v.Type,
		Params: v.Params,
	}
}

func (v StageJsonV1) mustConvert() Stage {
	return Stage{
		Id:   v.Id,
		Name: v.Name,
		NextStage: NextStage{
			Id:   v.Next,
			Node: nil,
		},
		Actions: v.mustConvertActions(),
	}
}

func (v StageJsonV1) mustConvertActions() Actions {
	a := Actions{}
	for _, aj := range v.Actions {
		a.Actions = append(a.Actions, aj.mustConvert())
	}
	return a
}

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
	"github.com/inflion/inflion/internal/ops/flow/context"
	"log"
)

type Expression struct {
	Input     string
	Operation string
	Value     string
}

type Conditions struct {
	Conditions []ConditionStage
}

func (c Conditions) Get(index int) ConditionStage {
	if index < len(c.Conditions) {
		return c.Conditions[index]
	}

	return ConditionStage{}
}

func (c Conditions) ReplaceById(id string, cond ConditionStage) Conditions {
	newConditions := Conditions{c.Conditions}
	for i, elm := range c.Conditions {
		if elm.Id == id {
			newConditions.Conditions[i] = cond
		} else {
			newConditions.Conditions[i] = elm
		}
	}
	return newConditions
}

type ConditionStage struct {
	Id          string
	Expressions []Expression
	IfTrue      NextStage
	IfFalse     NextStage
}

func (c ConditionStage) getId() string {
	return c.Id
}

func (c ConditionStage) isEnd() bool {
	return c.Id == endStageId
}

func (c ConditionStage) IsEmpty() bool {
	return c.Id == ""
}

func (c ConditionStage) isNormalStage() bool {
	return false
}

func (c ConditionStage) isConditionStage() bool {
	return true
}

func (c ConditionStage) Evaluate(ctx context.ExecutionContext) Stage {
	ex := c.Expressions[0]
	value := ctx.GetValueByPath(context.Path{Path: ex.Input})

	log.Printf("test: %s", value)

	if ex.Operation == "equals" {
		if value == ex.Value {
			return c.IfTrue.Stage
		} else {
			return c.IfFalse.Stage
		}
	}

	return nil
}

func (c ConditionStage) wasNextStageResolved() bool {
	return c.IfTrue.Stage != nil && c.IfFalse.Stage != nil
}

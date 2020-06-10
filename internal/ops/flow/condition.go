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

import "log"

type Expression struct {
	Input     string
	Operation string
	Value     string
}

type Conditions struct {
	Conditions []Condition
}

func (c Conditions) Get(index int) Condition {
	if index < len(c.Conditions) {
		return c.Conditions[index]
	}

	return Condition{}
}

func (c Conditions) ReplaceById(id string, cond Condition) Conditions {
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

type Condition struct {
	Id          string
	Expressions []Expression
	IfTrue      NextStage
	IfFalse     NextStage
}

func (c Condition) getId() string {
	return c.Id
}

func (c Condition) isEnd() bool {
	return c.Id == "__end__"
}

func (c Condition) IsEmpty() bool {
	return c.Id == ""
}

func (c Condition) isStage() bool {
	return false
}

func (c Condition) isCondition() bool {
	return true
}

func (c Condition) Evaluate(context ExecutionContext) Node {
	ex := c.Expressions[0]
	value := context.GetValueByPath(Path{Path: ex.Input})

	log.Printf("test: %s", value)

	if ex.Operation == "equals" {
		if value == ex.Value {
			return c.IfTrue.Node
		} else {
			return c.IfFalse.Node
		}
	}

	return nil
}

func (c Condition) wasNextStageResolved() bool {
	return c.IfTrue.Node != nil && c.IfFalse.Node != nil
}


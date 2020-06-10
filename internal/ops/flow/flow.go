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

type RecipeReader interface {
	Read() (Recipe, error)
}

type NextStage struct {
	Id   string
	Node Node
}

func (n NextStage) isEnd() bool {
	return n.Id == "__end__"
}

type Node interface {
	getId() string
	isEnd() bool
	isStage() bool
	isCondition() bool
	wasNextStageResolved() bool
}

type OpsFlow struct {
	reader RecipeReader
}

func NewOpsFlow(reader RecipeReader) *OpsFlow {
	return &OpsFlow{reader: reader}
}

func (o OpsFlow) Run(context ExecutionContext) (ExecutionResult, error) {
	log.Println("run flow")

	r, err := o.reader.Read()
	if err != nil {
		return ExecutionResult{}, err
	}

	r, err = r.LinkToNextStages()
	if err != nil {
		return ExecutionResult{}, err
	}

	return NewExecutor(AggregateActionLoader{}, context).run(r)
}

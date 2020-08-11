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

type Result struct {
	Message string
	Context context.ExecutionContext
}

type Executor struct {
	flow   Flow
	loader ActionLoader
}

func NewFlowExecutor(flow Flow, loader ActionLoader) *Executor {
	return &Executor{flow: flow, loader: loader}
}

func (e Executor) Run(ctx context.ExecutionContext) (Result, error) {
	c := e.exec(e.flow.Stages.getFirstStage(), ctx)

	return Result{
		Message: "success", // TODO handle errors. get from context.
		Context: c,
	}, nil
}

func (e Executor) exec(stage Stage, ctx context.ExecutionContext) context.ExecutionContext {
	if stage == nil { // this is probably the last stage (next stage is not specified)
		log.Println("stage is nil. end flow execution")
		return ctx
	}

	log.Printf("execute action: stage id: %s, context: %+v", stage.getId(), ctx)

	var nextStage Stage
	if stage.isNormalStage() {
		if s, ok := stage.(NormalStage); ok {
			ctx = e.execStage(s, ctx)

			if n := s.getNextStage().Stage; n != nil {
				nextStage = n
			} else {
				log.Printf("next stage is nil in stage at: %s", stage.getId())
			}
		} else {
			log.Println("stage is not a stage. it might be a bug.")
		}
	}

	if stage.isConditionStage() {
		if cond, ok := stage.(ConditionStage); ok {
			nextStage = cond.Evaluate(ctx)

			if nextStage == nil {
				log.Println("stage is nil. it might be a bug.")
			}
		} else {
			log.Println("stage is not a condition")
		}
	}

	if nextStage != nil {
		log.Printf("next stage: %s", nextStage.getId())
	} else {
		log.Printf("next stage is nil")
	}

	return e.exec(nextStage, ctx)
}

func (e Executor) execStage(stage NormalStage, ctx context.ExecutionContext) context.ExecutionContext {

	log.Printf("Run stage %s", stage.Id)
	log.Printf("Actions: %+v", stage.Actions)

	actionResults := ActionResults{}
	for _, a := range stage.Actions {
		ae, err := e.loader.Load(a)
		if err != nil {
			log.Println(err)
			return ctx
		}

		result, err := ae.Run(ctx)
		if err != nil {
			log.Println(err)
			return ctx
		}
		actionResults = append(actionResults, result)
	}
	ctx.AddField(stage.Name, actionResults)

	log.Printf("stage action results: %+v", actionResults)

	ctx.AddField("last", map[string]string{
		"status": actionResults.getExitMessage(),
	})

	return ctx
}

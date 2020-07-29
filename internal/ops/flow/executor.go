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
	"log"
)

type Result struct {
	Message string
	Context ExecutionContext
}

type Executor struct {
	flow   Flow
	loader ActionLoader
}

func NewFlowExecutor(flow Flow, loader ActionLoader) *Executor {
	return &Executor{flow: flow, loader: loader}
}

func (e Executor) Run(context ExecutionContext) (Result, error) {
	c := e.exec(e.flow.Stages.getFirstStage(), context)

	return Result{
		Message: "success", // TODO handle errors. get from context.
		Context: c,
	}, nil
}

func (e Executor) exec(stage Stage, context ExecutionContext) ExecutionContext {
	if stage == nil { // this is probably the last stage (next stage is not specified)
		log.Println("stage is nil. end flow execution")
		return context
	}

	log.Printf("execute action: stage id: %s, context: %+v", stage.getId(), context)

	var nextStage Stage
	if stage.isNormalStage() {
		if s, ok := stage.(NormalStage); ok {
			context = e.execStage(s, context)

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
			nextStage = cond.Evaluate(context)

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

	return e.exec(nextStage, context)
}

func (e Executor) execStage(stage NormalStage, context ExecutionContext) ExecutionContext {

	log.Printf("Run stage %s", stage.Id)
	log.Printf("Actions: %+v", stage.Actions)

	actionResults := ActionResults{}
	for _, a := range stage.Actions {
		ae, err := e.loader.Load(a)
		if err != nil {
			log.Println(err)
			return context
		}

		ar, err := ae.Run(context)
		if err != nil {
			log.Println(err)
			return context
		}
		actionResults = actionResults.append(ar)
	}

	log.Printf("action results: %+v", actionResults)

	context = e.addContextValuesFromActionResults(stage, actionResults, context)
	log.Printf("context: %+v", context.ExecutionFields)

	return context.AddFields("last", ExecutionFields{
		Values: map[string]interface{}{
			"status": actionResults.getExitMessage(),
		},
	})
}

func (e Executor) addContextValuesFromActionResults(stage NormalStage, results ActionResults, context ExecutionContext) ExecutionContext {
	for _, r := range results {
		values := map[string]interface{}{}
		for k, v := range r.Outputs {
			values[k] = v
		}
		fields := ExecutionFields{
			Values: values,
		}
		context = context.AddFields(stage.Name, fields)
	}
	return context
}

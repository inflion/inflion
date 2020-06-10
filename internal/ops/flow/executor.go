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

type Executor interface {
	run(r Recipe) (ExecutionResult, error)
}

type ExecutionResult struct {
	Message string
	Context ExecutionContext
}

type DefaultExecutor struct {
	loader  ActionLoader
	context ExecutionContext
}

func NewExecutor(loader ActionLoader, context ExecutionContext) Executor {
	return DefaultExecutor{loader: loader, context: context}
}

func (e DefaultExecutor) run(r Recipe) (ExecutionResult, error) {
	c := e.exec(r.Stages.getFirstStage(), e.context)

	return ExecutionResult{
		Message: "success", // TODO handle errors. get from context.
		Context: c,
	}, nil
}

func (e DefaultExecutor) exec(node Node, context ExecutionContext) ExecutionContext {
	if node == nil {
		log.Println("node is nil")
		return context
	}

	log.Printf("execute action: node id: %s, context: %+v", node.getId(), context)
	var nextNode Node

	if node.isStage() {
		if s, ok := node.(Stage); ok {
			context = e.execStage(s, context)

			if n := s.getNextNode(); n != nil {
				nextNode = n
			} else {
				log.Printf("next node is nil in stage at: %s", node.getId())
			}
		} else {
			log.Println("node is not a stage. it might be a bug.")
		}
	}

	if node.isCondition() {
		if cond, ok := node.(Condition); ok {
			nextNode = cond.Evaluate(context)

			if nextNode == nil {
				log.Println("node is nil. it might be a bug.")
			}
		} else {
			log.Println("node is not a condition")
		}
	}

	return e.exec(nextNode, context)
}

func (e DefaultExecutor) execStage(stage Stage, context ExecutionContext) ExecutionContext {
	actionResults := actionResults{}

	log.Printf("actions: %+v", stage.Actions.Actions)

	for _, a := range stage.Actions.Actions {
		ae, err := e.loader.Load(a)
		if err != nil {
			log.Println(err)
			return context
		}

		ar, err := ae.Run(context, a)
		if err != nil {
			log.Println(err)
			return context
		}
		actionResults = actionResults.append(ar)
	}

	log.Printf("action results: %+v", actionResults.results)

	context = e.addContextValuesFromActionResults(stage, actionResults, context)
	log.Printf("context: %+v", context.ExecutionFields)

	return context.addFields("last", ExecutionFields{
		Values: map[string]interface{}{
			"status": actionResults.getExitMessage(),
		},
	})
}

func (e DefaultExecutor) addContextValuesFromActionResults(stage Stage, results actionResults, context ExecutionContext) ExecutionContext {
	for _, r := range results.results {
		values := map[string]interface{}{}
		for k, v := range r.Outputs {
			values[k] = v
		}
		fields := ExecutionFields{
			Values: values,
		}
		context = context.addFields(stage.Name, fields)
	}
	return context
}

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
	"github.com/inflion/inflion/internal/ops/flow/action"
	"log"
)

type ActionResult struct {
	Action     Action
	Outputs    map[string]string
	ExitStatus bool
}

type ActionExecutor interface {
	Run(context ExecutionContext, action Action) (ActionResult, error)
}

type NullActionExecutor struct {
}

func (n NullActionExecutor) Run(_ ExecutionContext, _ Action) (ActionResult, error) {
	return ActionResult{}, nil
}

type MatcherActionExecutor struct {
}

func (m MatcherActionExecutor) Run(_ ExecutionContext, action Action) (ActionResult, error) {
	log.Println("execute action: " + action.Type)
	return ActionResult{
		Action: action,
		Outputs: map[string]string{
			"result": "false",
		},
		ExitStatus: false,
	}, nil
}

type InstanceActionExecutor struct {
}

func (i InstanceActionExecutor) Run(_ ExecutionContext, action Action) (ActionResult, error) {
	log.Println("execute action: " + action.Type)
	log.Printf("action params: %+v", action.Params)
	return ActionResult{
		Action: action,
		Outputs: map[string]string{
			"result": "true",
		},
		ExitStatus: true,
	}, nil
}

type NotificationActionExecutor struct {
}

func (n NotificationActionExecutor) Run(ec ExecutionContext, a Action) (ActionResult, error) {
	log.Println("execute action: " + a.Type)

	ec.GetValueByPath(NewPath("trigger.event"))

	notifier := action.SlackNotifier{}
	notifier.Notify(a.Params)

	return ActionResult{
		Action: a,
		Outputs: map[string]string{
			"result": "true",
		},
		ExitStatus: true,
	}, nil
}

type LoggingActionExecutor struct {
}

func (l LoggingActionExecutor) Run(e ExecutionContext, action Action) (ActionResult, error) {
	log.Printf("log: %s", action.Params)

	return ActionResult{
		Action: action,
		Outputs: map[string]string{
			"result": "false",
		},
		ExitStatus: true,
	}, nil
}

type ParamsActionExecutor struct {
}

func (p ParamsActionExecutor) Run(e ExecutionContext, a Action) (ActionResult, error) {
	log.Printf("log: %s", a.Params)

	outputs := map[string]string{}

	for k, param := range a.Params {
		outputs[k] = param
	}

	return ActionResult{
		Action: a,
		Outputs: outputs,
		ExitStatus: true,
	}, nil
}

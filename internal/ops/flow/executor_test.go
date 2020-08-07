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
	"testing"
)

type MockActionExecutor struct {
	action Action
}

func (m MockActionExecutor) Run(_ ExecutionContext) (ActionResult, error) {
	log.Println("exec mock action")

	return ActionResult{
		Action: m.action,
		Outputs: map[string]string{
			"output1-key": "output1-value",
			"output2-key": "output2-value",
		},
		ExitStatus: true,
	}, nil
}

type MockActionLoader struct{}

func (m MockActionLoader) Load(action Action) (ActionExecutor, error) {
	return MockActionExecutor{action: action}, nil
}

func Test_last_status_should_be_success(t *testing.T) {
	flow, _ := MockOpsFlow{}.Read()

	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	result, _ := NewFlowExecutor(r, NewAggregateActionLoader()).Run(NewExecutionContext())

	got := result.Context.GetValueByPath(NewPath("last.status"))
	want := "success"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func Test_stage1_result_should_have_message(t *testing.T) {
	flow, _ := MockOpsFlow{}.Read()

	linkedFlow, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	result, _ := NewFlowExecutor(linkedFlow, MockActionLoader{}).Run(NewExecutionContext())

	got := result.Context.GetValueByPath(NewPath("stage1-name.output1-key"))
	want := "output1-value"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

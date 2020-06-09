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

type MockActionLoader struct {
}

type MockActionExecutor struct {
}

func (m MockActionExecutor) Run(_ ExecutionContext, action Action) (ActionResult, error) {
	log.Println("exec mock action")

	return ActionResult{
		Action: action,
		Outputs: map[string]string{
			"output1-key": "output1-value",
			"output2-key": "output2-value",
		},
		ExitStatus: true,
	}, nil
}

func (m MockActionLoader) Load(_ Action) (ActionExecutor, error) {
	return MockActionExecutor{}, nil
}

func Test_last_status_should_be_success(t *testing.T) {
	recipe, _ := MockRecipeReader{}.Read()

	r, err := recipe.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	result, _ := NewExecutor(MockActionLoader{}, NewExecutionContext()).run(r)

	got := result.Context.GetValueByPath(NewPath("last.status"))
	want := "success"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func Test_stage1_result_should_have_message(t *testing.T) {
	recipe, _ := MockRecipeReader{}.Read()

	r, err := recipe.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	result, _ := NewExecutor(MockActionLoader{}, NewExecutionContext()).run(r)

	got := result.Context.GetValueByPath(NewPath("stage1-name.output1-key"))
	want := "output1-value"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

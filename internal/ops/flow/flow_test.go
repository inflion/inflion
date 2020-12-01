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
	inflionEvent "github.com/inflion/inflion/internal/ops/event"
	"github.com/inflion/inflion/internal/ops/flow/context"
	"io/ioutil"
	"testing"
)

var flow Flow

func init() {
	flow, _ = MockOpsFlow{}.Read()
}

func TestFlow(t *testing.T) {
	rawEvent, err := ioutil.ReadFile("event.json")
	if err != nil {
		t.Error(err)
	}

	event, err := inflionEvent.NewInflionEvent("test", rawEvent)
	if err != nil {
		t.Error(err)
	}
	ec := context.NewExecutionContextWithEvent(event)

	fe := NewFlowExecutor(flow, NewAggregateActionLoader())
	_, _ = fe.Run(ec)
}

func Test_next_stage_should_be_cond1_via_first_stage_then_next_stage(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	want := "cond1-id"
	got := r.Stages.getFirstStage().getNextStage().Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

func Test_next_stage_should_be_cond1_in_stage1(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	want := "cond1-id"
	got := r.Stages.get(0).getNextStage().Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

func Test_next_stage_should_be_end_in_stage4(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	want := "__end__"
	got := r.Stages.get(3).getNextStage().Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

func Test_next_stage_should_be_stage2_if_true_in_condition_one(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	want := "stage2-id"
	got := r.Conditions.Get(0).IfTrue.Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

func Test_next_stage_should_be_stage3_if_false_in_condition_one(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	want := "stage3-id"
	got := r.Conditions.Get(0).IfFalse.Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

func Test_non_existent_condition_should_be_empty(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	got := r.Conditions.Get(999).IsEmpty()
	if true != got {
		t.Errorf("got %v want true", got)
	}
}

// Testing boundary value
func Test_condition1_should_not_be_empty(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	got := r.Conditions.Get(0).IsEmpty()
	if false != got {
		t.Errorf("got %v want false", got)
	}
}

func Test_next_stage_of_stage1_should_be_condition1(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	want := "cond1-id"
	got := r.Stages.get(0).NextStage.Id
	if want != got {
		t.Errorf("got %s want %s", want, got)
	}
}

// Testing boundary value
func Test_non_existent_stage_should_be_empty(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	got := r.Stages.get(999).isEmpty()
	if true != got {
		t.Errorf("got %v want true", got)
	}
}

// Testing boundary value
func Test_stage_should_be_exists(t *testing.T) {
	r, err := flow.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	got := r.Stages.get(2).isEmpty()
	if false != got {
		t.Errorf("got %v want false", got)
	}
}

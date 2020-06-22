// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package broker

import (
	"github.com/inflion/inflion/internal/ops/flow/store"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/ops/rule"
	"io/ioutil"
	"testing"
)

type mockEventMatcher struct {
}

func (m mockEventMatcher) GetRulesMatchesTo(event monitor.MonitoringEvent) ([]rule.Rule, error) {
	return []rule.Rule{
		{
			RuleName: "rule1",
			Target:   "a9aebdfd-4d43-42b7-9315-b1b89e03ffa4",
			Conditions: rule.Conditions{
				Conditions: []rule.Condition{
					{
						TargetAttr: "target_attribute",
						Matcher: &rule.ContainsMatcher{
							Value: "error",
						},
					},
				},
			},
		},
	}, nil
}

type mockStore struct {
	store.Store
}

func (e mockStore) Get(request store.FlowGetRequest) (store.FlowGetResponse, error) {
	jsonForTest := "../flow/sample_flow.json"

	bytes, err := ioutil.ReadFile(jsonForTest)
	if err != nil {
		panic(err)
	}
	return store.FlowGetResponse{
		Body: string(bytes),
	}, nil
}

func TestEventProcessor(t *testing.T) {
	e := NewEventProcessor(mockStore{}, mockEventMatcher{})
	err := e.process(
		monitor.MonitoringEvent{
			Project: "sandbox",
			Body:    map[string]interface{}{"type": "test", "message": "message"},
		},
	)
	if err != nil {
		t.Errorf("got an error %+v", err)
	}
}

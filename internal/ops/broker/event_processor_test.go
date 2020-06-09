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
	"context"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/ops/rule"
	"github.com/inflion/inflion/internal/store"
	"io/ioutil"
	"testing"
	"time"
)

type mockEventMatcher struct {
}

func (m mockEventMatcher) getRulesMatchesTo(event monitor.MonitoringEvent) ([]rule.Rule, error) {
	return []rule.Rule{
		{
			RuleName: "rule1",
			Target:   "test",
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

type mockQuerier struct {
	store.Querier
}

func (m mockQuerier) GetFlowByName(ctx context.Context, arg store.GetFlowByNameParams) ([]store.Flow, error) {
	jsonForTest := "../flow/json/sample_flow.json"

	bytes, err := ioutil.ReadFile(jsonForTest)
	if err != nil {
		panic(err)
	}
	return []store.Flow{
		{
			ID:        1,
			ProjectID: 1,
			FlowName:  "flow",
			Body:      bytes,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}, nil
}

func TestEventProcessor(t *testing.T) {
	e := newDefaultEventProcessor(mockQuerier{}, mockEventMatcher{})
	err := e.process(
		monitor.MonitoringEvent{
			Type:      "mock-event",
			ProjectId: 1,
			Message:   "message",
			Values:    map[string]interface{}{"test": "test"},
		},
	)
	if err != nil {
		t.Errorf("got an error %+v", err)
	}
}

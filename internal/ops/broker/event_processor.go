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
	"github.com/google/uuid"
	"github.com/inflion/inflion/internal/ops/flow"
	"github.com/inflion/inflion/internal/ops/flow/store"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/ops/rule"
	"log"
)

type eventProcessor interface {
	process(event monitor.MonitoringEvent) error
}

type defaultEventProcessor struct {
	matcher rule.EventMatcher
	store   store.Store
}

func NewEventProcessor(store store.Store, matcher rule.EventMatcher) eventProcessor {
	return defaultEventProcessor{
		store:   store,
		matcher: matcher,
	}
}

func (d defaultEventProcessor) process(event monitor.MonitoringEvent) error {
	matchedRules, err := d.matcher.GetRulesMatchesTo(event)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, matchedRule := range matchedRules {
		flowId, err := uuid.Parse(matchedRule.Target)
		if err != nil {
			return err
		}

		r, err := d.store.Get(
			store.FlowGetRequest{
				Id:      flowId,
				Project: event.Project,
			},
		)
		if err != nil {
			log.Println(err)
			return err
		}

		f := flow.NewOpsFlow(ByteRecipeReader{body: []byte(r.Body)})
		result, err := f.Run(flow.NewExecutionContext())
		if err != nil {
			log.Println(err)
		}
		log.Printf("flow execution result: %+v", result)
	}

	return nil
}

type ByteRecipeReader struct {
	body []byte
}

func (b ByteRecipeReader) Read() (flow.Recipe, error) {
	recipe, err := flow.Unmarshal(b.body)
	if err != nil {
		return flow.Recipe{}, err
	}
	return recipe, nil
}

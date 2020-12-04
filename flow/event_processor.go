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
	"github.com/google/uuid"
	"github.com/inflion/inflion/flow/context"
	"github.com/inflion/inflion/flow/event"
	"github.com/inflion/inflion/flow/matcher"
	"log"
)

type EventProcessor interface {
	Process(event.InflionEvent) error
}

type defaultEventProcessor struct {
	matcher matcher.EventMatcher
	store   Store
}

func NewEventProcessor(store Store, matcher matcher.EventMatcher) EventProcessor {
	return defaultEventProcessor{
		store:   store,
		matcher: matcher,
	}
}

func (d defaultEventProcessor) Process(event event.InflionEvent) error {
	matchedRules, err := d.matcher.GetRulesMatchesTo(event)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("matched rules: %+v", matchedRules)

	for _, matchedRule := range matchedRules {
		flowId, err := uuid.Parse(matchedRule.Target)
		if err != nil {
			return err
		}

		resp, err := d.store.Get(
			FlowGetRequest{
				Id:      flowId,
				Project: event.Project(),
			},
		)
		if err != nil {
			log.Println(err)
			return err
		}

		f, err := ByteFlowReader{body: []byte(resp.Body)}.Read()
		if err != nil {
			log.Println(err)
		}

		result, err := NewFlowExecutor(f, NewAggregateActionLoader()).Run(context.NewExecutionContextWithEvent(&event))

		if err != nil {
			log.Println(err)
		}
		log.Printf("flow execution result: %+v", result)
	}

	return nil
}

// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package processor

import (
	"github.com/google/uuid"
	inflionEvent "github.com/inflion/inflion/internal/ops/event"
	"github.com/inflion/inflion/internal/ops/flow"
	"github.com/inflion/inflion/internal/ops/flow/context"
	"github.com/inflion/inflion/internal/ops/flow/store"
	"github.com/inflion/inflion/internal/ops/rule"
	"log"
)

type EventProcessor interface {
	Process(inflionEvent.InflionEvent) error
}

type defaultEventProcessor struct {
	matcher rule.EventMatcher
	store   store.Store
}

func NewEventProcessor(store store.Store, matcher rule.EventMatcher) EventProcessor {
	return defaultEventProcessor{
		store:   store,
		matcher: matcher,
	}
}

func (d defaultEventProcessor) Process(event inflionEvent.InflionEvent) error {
	log.Printf("processing evnet: %+v", event)

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
			store.FlowGetRequest{
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

		result, err := flow.NewFlowExecutor(f, flow.NewAggregateActionLoader()).Run(context.NewExecutionContextWithEvent(&event))

		if err != nil {
			log.Println(err)
		}
		log.Printf("flow execution result: %+v", result)
	}

	return nil
}

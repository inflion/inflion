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

type EventProcessor interface {
	process(event monitor.MonitoringEvent) error
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

func (d defaultEventProcessor) process(event monitor.MonitoringEvent) error {
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
				Project: event.Project,
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

		ec := flow.NewExecutionContext()
		ec.AddFields("project", flow.ExecutionFields{
			Values: map[string]interface{}{"id": event.Project},
		})
		ec.AddFields("event", flow.ExecutionFields{
			Values: event.Body,
		})
		ec.AddFields("raw-event", flow.ExecutionFields{
			Values: map[string]interface{}{"json": event.RawBody},
		})

		result, err := flow.NewFlowExecutor(f, flow.NewAggregateActionLoader()).Run(ec)
		if err != nil {
			log.Println(err)
		}
		log.Printf("flow execution result: %+v", result)
	}

	return nil
}

type ByteFlowReader struct {
	body []byte
}

func (b ByteFlowReader) Read() (flow.Flow, error) {
	recipe, err := flow.Unmarshal(b.body)
	if err != nil {
		return flow.Flow{}, err
	}
	return recipe, nil
}

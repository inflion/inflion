// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package rule

import (
	"github.com/inflion/inflion/internal/ops/monitor"
	"log"
)

type EventMatcher interface {
	GetRulesMatchesTo(event monitor.MonitoringEvent) ([]Rule, error)
}

func NewEventMatcher(store Store) EventMatcher {
	return &storeEventMatcher{store: store}
}

type storeEventMatcher struct {
	store Store
}

func (s *storeEventMatcher) GetRulesMatchesTo(event monitor.MonitoringEvent) ([]Rule, error) {
	var matchedRules []Rule

	rules, err := s.store.GetRules(event.Project)
	if err != nil {
		return matchedRules, err
	}

	log.Printf("rules: %+v", rules)

	for _, rule := range rules {
		if rule.Conditions.match(event) {
			matchedRules = append(matchedRules, rule)
		}
	}

	return matchedRules, nil
}

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
	//store2 "github.com/inflion/inflion/internal/ops/rule/store"
)

type EventMatcher interface {
	GetRulesMatchesTo(event monitor.MonitoringEvent) ([]Rule, error)
}

type querierEventMatcher struct {
	//store store2.Store
}

func (n *querierEventMatcher) GetRulesMatchesTo(event monitor.MonitoringEvent) ([]Rule, error) {
	var matchedRules []Rule

	// TODO use interface instead of store

	//rules, err := n.store.GetRules(event.ProjectId)
	//if err != nil {
	//	return matchedRules, err
	//}

	//for _, rule := range rules {
		//if rule.Conditions.match(event) {
		//	matchedRules = append(matchedRules, rule)
		//}
	//}

	return matchedRules, nil
}

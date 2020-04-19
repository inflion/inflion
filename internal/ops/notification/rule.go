// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package notification

import (
	"context"
	"encoding/json"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/store"
	"log"
)

type ruleMatcher interface {
	match(value interface{}) bool
}

type notificationRule struct {
	querier store.Querier
}

type Rules struct {
	Rules []MatchRule
}

type Rule struct {
	Name  string
	Rules Rules
}

type JsonMatchRule struct {
	EventType  string `json:"event_type"`
	TargetAttr string `json:"target_attr"`
	Type       string `json:"match_type"`
	Value      string `json:"match_value"`
}

type MatchRule struct {
	name        string
	attribute   string
	ruleMatcher ruleMatcher
}

func (r *Rules) match(event monitor.MonitoringEvent) bool {
	var matches []bool
	for _, rule := range r.Rules {
		matches = append(matches, rule.match(event))
	}

	result := true
	for _, b := range matches {
		if b == false {
			result = false
			break
		}
	}
	return result
}

func newRules(ruleName string, jsonMatchRules []JsonMatchRule) Rules {
	var rules []MatchRule

	for _, jsonMatchRule := range jsonMatchRules {
		rules = append(rules, MatchRule{
			name:        ruleName,
			attribute:   jsonMatchRule.TargetAttr,
			ruleMatcher: &ruleMatcherContains{jsonMatchRule.Value},
		})
	}

	return Rules{
		Rules: rules,
	}
}

func (m *MatchRule) match(event monitor.MonitoringEvent) bool {
	if value, ok := event.GetValue(m.attribute); ok {
		return m.ruleMatcher.match(value)
	}
	return false
}

func (n *notificationRule) match(event monitor.MonitoringEvent) ([]Rule, error) {
	var matchedRules []Rule

	rules, err := n.getRules(event.ProjectId)
	if err != nil {
		return matchedRules, err
	}

	for _, rule := range rules {
		if rule.Rules.match(event) {
			matchedRules = append(matchedRules, rule)
		}
	}

	return matchedRules, nil
}

func (n *notificationRule) getRules(projectId int64) ([]Rule, error) {
	ctx := context.Background()

	var rules []Rule

	dbRules, err := n.querier.GetNotificationRules(ctx, projectId)
	if err != nil {
		log.Println(err)
		return []Rule{}, err
	}

	for _, dbRule := range dbRules {
		var jsonMatchRules []JsonMatchRule
		err := json.Unmarshal(dbRule.Rules, &jsonMatchRules)
		if err != nil {
			log.Println(err)
			return []Rule{}, err
		}

		rules = append(rules, Rule{
			Name:  dbRule.RuleName,
			Rules: newRules(dbRule.RuleName, jsonMatchRules),
		})
	}

	return rules, nil
}

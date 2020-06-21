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

type condMatcher interface {
	match(value interface{}) bool
}

type Condition struct {
	TargetAttr string
	MatchType  string
	MatchValue string
	Matcher    condMatcher
}

func (m *Condition) match(event monitor.MonitoringEvent) bool {
	if m.MatchType == "contains" {
		m.Matcher = &ContainsMatcher{Value: m.MatchValue}
	} else if m.MatchType == "exact" {
		m.Matcher = &exactMatcher{value: m.MatchValue}
	} else {
		log.Printf("no such match type: %s", m.MatchType)
		return false
	}

	if eventValue, ok := event.GetValue(m.TargetAttr); ok {
		return m.Matcher.match(eventValue)
	}

	return false
}

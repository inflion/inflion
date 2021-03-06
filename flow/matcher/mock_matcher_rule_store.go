// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package matcher

type MockMatcherRuleStore struct{}

func (m MockMatcherRuleStore) GetRules() ([]Rule, error) {
	return []Rule{{
		RuleName: "",
		Target:   "",
		Conditions: Conditions{Conditions: []Condition{{
			TargetAttr: "",
			Matcher:    nil,
		}}},
	}}, nil
}

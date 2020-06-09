// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package jsonv1

import (
	"encoding/json"
	"github.com/inflion/inflion/internal/ops/rule"
)

func UnmarshalV1(v1FormattedJson []byte) (*RuleRootJsonV1, error) {
	v1 := RuleRootJsonV1{}
	err := json.Unmarshal(v1FormattedJson, &v1)
	if err != nil {
		return nil, err
	}
	return &v1, nil
}

type RuleRootJsonV1 struct {
	Metadata MetadataJson `json:"metadata"`
	Body     struct {
		Name       string            `json:"name"`
		Target     string            `json:"target"`
		Conditions []ConditionJsonV1 `json:"conditions"`
	} `json:"body"`
}

type ConditionJsonV1 struct {
	TargetAttr string `json:"target_attr"`
	MatchType  string `json:"match_type"`
	MatchValue string `json:"match_value"`
}

func (r RuleRootJsonV1) mustConvertConditions() []rule.Condition {
	var c []rule.Condition
	for _, jc := range r.Body.Conditions {
		c = append(c, rule.Condition{
			TargetAttr: jc.TargetAttr,
			MatchType:  jc.MatchType,
			MatchValue: jc.MatchValue,
		})
	}
	return c
}

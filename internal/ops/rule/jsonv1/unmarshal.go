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
	"fmt"
	"github.com/inflion/inflion/internal/ops/rule"
)

func Unmarshal(rawJson []byte) (rule.Rule, error) {
	m := MetadataJson{}
	err := json.Unmarshal(rawJson, &m)
	if err != nil {
		return rule.Rule{}, err
	}

	if m.Metadata.Format.Version == 1 {
		v1, err := UnmarshalV1(rawJson)
		if err != nil {
			return rule.Rule{}, err
		}

		return rule.Rule{
			RuleName:   v1.Body.Name,
			Target:     v1.Body.Target,
			Conditions: rule.Conditions{Conditions: v1.mustConvertConditions()},
		}, nil
	}

	return rule.Rule{}, fmt.Errorf("json version %d not supported", m.Metadata.Format.Version)
}

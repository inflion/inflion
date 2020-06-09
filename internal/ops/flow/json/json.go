// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package json

import (
	"encoding/json"
	"fmt"
	"github.com/inflion/inflion/internal/ops/flow"
)

func Unmarshal(rawJson []byte) (flow.Recipe, error) {
	m := MetadataJson{}
	err := json.Unmarshal(rawJson, &m)
	if err != nil {
		return flow.Recipe{}, err
	}

	if m.Metadata.Format.Version == 1 {
		v1, err := UnmarshalV1(rawJson)
		if err != nil {
			return flow.Recipe{}, err
		}
		c := flow.Conditions{}
		for _, cj := range v1.Body.Conditions {
			c.Conditions = append(c.Conditions, cj.mustConvert())
		}

		s := flow.Stages{}
		for _, sj := range v1.Body.Stages {
			s.Stages = append(s.Stages, sj.mustConvert())
		}

		return flow.Recipe{
			Conditions: c,
			Stages:     s,
		}, nil
	}

	return flow.Recipe{}, fmt.Errorf("json version %d not supported", m.Metadata.Format.Version)
}

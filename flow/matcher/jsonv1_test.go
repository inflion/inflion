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

import (
	"io/ioutil"
	"testing"
)

func TestUnmarshalV1(t *testing.T) {
	jsonForTest := "v1.json"

	bytes, err := ioutil.ReadFile(jsonForTest)
	if err != nil {
		t.Error(err)
	}

	rule, err := Unmarshal(bytes)
	if err != nil {
		t.Error(err)
	}

	if rule.Conditions.Conditions[0].TargetAttr != "test" {
		t.Error("probably unmarshal failed")
	}
}

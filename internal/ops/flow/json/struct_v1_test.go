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
	"io/ioutil"
	"testing"
)

func TestFlowStruct(t *testing.T) {
	jsonForTest := "sample_flow.json"

	bytes, err := ioutil.ReadFile(jsonForTest)
	if err != nil {
		t.Error(err)
	}

	r, err := Unmarshal(bytes)
	if err != nil {
		t.Error(err)
	}

	if r.Conditions.Conditions[0].IfTrue.Id != "stage_id_3" {
		t.Errorf("First conditions not have id 3, got: %s", r.Conditions.Conditions[0].IfTrue.Id)
	}
}

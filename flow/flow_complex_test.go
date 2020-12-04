// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package flow

import (
	"io/ioutil"
	"log"
	"testing"
)

func read() Flow {
	bytes, err := ioutil.ReadFile("./test.json")
	if err != nil {
		log.Fatal(err)
	}

	recipe, err := Unmarshal(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return recipe
}

func Test_complex(t *testing.T) {
	f := read()
	r, err := f.LinkToNextStages()
	if err != nil {
		t.Error(err)
	}

	first := r.Stages.getFirstStage()

	if first.Id != "params" {
		t.Error("first stage must be params")
	}

	if first.NextStage.Stage.getId() != "config" {
		// TODO BUG FIX
		t.Errorf("second stage must be config. actual: %s", first.NextStage.Stage.getId())
	}
}

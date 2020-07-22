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
	"encoding/json"
	"github.com/inflion/inflion/internal/ops/monitor"
	"io/ioutil"
	"testing"
)

func TestFlow(t *testing.T) {
	jsonForTest := "event.json"

	bytes, err := ioutil.ReadFile(jsonForTest)
	if err != nil {
		t.Error(err)
	}

	event := monitor.MonitoringEvent{}

	err = json.Unmarshal(bytes, &event)
	if err != nil {
		t.Error(err)
	}

	event.RawEventBody = bytes

	ec := NewExecutionContext()
	ec.AddFields("project", ExecutionFields{
		Values: map[string]interface{}{"id": event.Project},
	})
	ec.AddFields("event", ExecutionFields{
		Values: event.EventBody,
	})
	ec.AddFields("raw-event", ExecutionFields{
		Values: map[string]interface{}{"json": event.RawEventBody},
	})

	flow := NewOpsFlow(MockRecipeReader{})
	_, err = flow.Run(ec)
	if err != nil {
		t.Error(err)
	}
}

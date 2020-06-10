// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package broker

import (
	"github.com/inflion/inflion/internal/ops/monitor"
	"testing"
)

type mockEventProcessor struct {
	processedEventCount int
}

func (m *mockEventProcessor) process(event monitor.MonitoringEvent) error {
	m.processedEventCount = m.processedEventCount + 1
	return nil
}

type mockConsumer struct {
}

func (m mockConsumer) consume(processor eventProcessor) {
	_ = processor.process(monitor.MonitoringEvent{
		Type:      "mock-event",
		ProjectId: 1,
		Message:   "message",
		Values:    map[string]interface{}{"test": "test"},
	})
}

func TestBroker(t *testing.T) {
	mep := mockEventProcessor{}
	b := NewBroker(mockConsumer{}, &mep)
	b.Run()

	if mep.processedEventCount != 1 {
		t.Error("event processed count not equal to 1")
	}
}
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
	"encoding/json"
	inflionEvent "github.com/inflion/inflion/internal/ops/event"
	"github.com/inflion/inflion/internal/ops/processor"
	"log"
	"testing"
)

type mockEventProcessor struct {
	processedEventCount int
}

func (m *mockEventProcessor) Process(_ inflionEvent.InflionEvent) error {
	m.processedEventCount = m.processedEventCount + 1
	return nil
}

type mockConsumer struct{}

func (m mockConsumer) consume(processor processor.EventProcessor) {
	bytes, err := json.Marshal(map[string]interface{}{"type": "test", "message": "message"})
	if err != nil {
		log.Print(err)
	}
	ie, err := inflionEvent.NewInflionEvent("sandbox", bytes)
	if err != nil {
		log.Print(err)
	}
	_ = processor.Process(*ie)
}

func TestBroker(t *testing.T) {
	mep := mockEventProcessor{}
	b := NewBroker(mockConsumer{}, &mep)
	b.Run()

	if mep.processedEventCount != 1 {
		t.Error("event processed count not equal to 1")
	}
}

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
	_ "github.com/lib/pq"
)

const topicName = "monitoring-events"

type consumer interface {
	consume(processor eventProcessor)
}

// Broker is to handle monitoring events.
type Broker struct {
	consumer consumer
	processor eventProcessor
}

func NewBroker(consumer consumer, processor eventProcessor) Broker {
	return Broker{
		consumer: consumer,
		processor: processor,
	}
}

func (b *Broker) Run() {
	b.consumer.consume(b.processor)
}

// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package metrics

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type producer struct {
	producer *nsq.Producer
}

func newProducer(config *nsq.Config) *producer {
	p, err := nsq.NewProducer("nsqd:4150", config)
	if err != nil {
		panic(err)
	}
	return &producer{
		producer: p,
	}
}

func (p *producer) produce(body []byte) (err error) {
	topicName := "metrics"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = p.producer.Publish(topicName, body)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func (p *producer) Stop() {
	p.Stop()
}

// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package producer

import (
	"github.com/nsqio/go-nsq"
	"log"
)

const topicName = "monitoring-events"
const nsqdAddress = "nsqd:4150"

type Producer struct {
	producer *nsq.Producer
}

func NewProducer() *Producer {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(nsqdAddress, config)
	if err != nil {
		panic(err)
	}
	return &Producer{
		producer: p,
	}
}

func (p *Producer) Produce(event []byte) error {
	err := p.producer.Publish(topicName, event)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *Producer) Stop() {
	p.Stop()
}

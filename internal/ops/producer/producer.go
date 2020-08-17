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
	"encoding/json"
	inflionEvent "github.com/inflion/inflion/internal/ops/event"
	"github.com/nsqio/go-nsq"
	"log"
)

const topicName = "monitoring-events"
const nsqdAddress = "nsqd:4150"

var producer *nsq.Producer

type Producer struct{}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) Produce(event inflionEvent.InflionEvent) (err error) {
	if producer == nil {
		producer, err = nsq.NewProducer(nsqdAddress, nsq.NewConfig())
		if err != nil {
			log.Println(err)
			return err
		}
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
		return err
	}

	err = producer.Publish(topicName, message)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (p *Producer) Stop() {
	p.Stop()
}

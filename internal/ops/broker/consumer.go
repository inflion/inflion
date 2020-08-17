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
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

type nsqConsumer struct {
	processor processor.EventProcessor
}

func NewNsqConsumer() Consumer {
	return &nsqConsumer{}
}

func (n *nsqConsumer) consume(processor processor.EventProcessor) {
	n.processor = processor
	nsqlookupdHost := os.Getenv("NSQLOOKUPD_HOST")
	nsqlookupdPort := os.Getenv("NSQLOOKUPD_PORT")

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topicName, "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(n)

	err = consumer.ConnectToNSQLookupd(nsqlookupdHost + ":" + nsqlookupdPort)
	if err != nil {
		log.Fatal(err)
	}

	<-consumer.StopChan
	consumer.Stop()
}

func (n *nsqConsumer) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}

	e := inflionEvent.InflionEvent{}
	err := json.Unmarshal(message.Body, &e)
	if err != nil {
		log.Println(err)
		return err
	}

	err = n.processor.Process(e)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

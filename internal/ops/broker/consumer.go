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
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

type nsqConsumer struct {
	processor eventProcessor
}

func (n *nsqConsumer) consume(processor eventProcessor) {
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

	event := monitor.MonitoringEvent{}
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = n.processor.process(event)
	if err != nil {
		log.Println(err)
		return nil
	}

	return nil
}
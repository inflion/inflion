// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package notification

import (
	"context"
	"encoding/json"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/store"
	_ "github.com/lib/pq"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

const topicName = "monitoring-events"

// Broker is to handle monitoring events.
type Broker struct {
	querier store.Querier
}

type messageHandler struct {
	querier store.Querier
}

func NewBroker(querier store.Querier) Broker {
	staticActionRegistry.actions["slack"] = slackNotificationAction{notifier: newThrottledNotification(querier, newNotification(querier).Notifiers)}

	return Broker{
		querier: querier,
	}
}

func (b *Broker) Run() {
	nsqlookupdHost := os.Getenv("NSQLOOKUPD_HOST")
	nsqlookupdPort := os.Getenv("NSQLOOKUPD_PORT")

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topicName, "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&messageHandler{
		querier: b.querier,
	})

	err = consumer.ConnectToNSQLookupd(nsqlookupdHost + ":" + nsqlookupdPort)
	if err != nil {
		log.Fatal(err)
	}

	<-consumer.StopChan
	consumer.Stop()
}

func (h *messageHandler) HandleMessage(message *nsq.Message) error {
	ctx := context.Background()

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

	notificationRule := notificationRule{querier: h.querier}
	matchedRules, err := notificationRule.match(event)
	if err != nil {
		log.Println(err)
		return err
	}

	actions, err := h.querier.GetActions(ctx, event.ProjectId)
	if err != nil {
		log.Println(err)
	}

	for _, rule := range matchedRules {
		for _, rawAction := range actions {
			ae, err := newActionExecutor(rawAction, event)
			if err != nil {
				log.Println(err)
				continue
			}
			ae.exec(rule)
		}
	}

	return nil
}

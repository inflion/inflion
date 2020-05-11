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
	"context"
	"database/sql"
	"encoding/json"
	"github.com/inflion/inflion/internal/timescale"
	_ "github.com/lib/pq"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

type consumer struct {
	querier timescale.Querier
}

type messageHandler struct {
	query timescale.Querier
}

func newConsumer(querier timescale.Querier) consumer {
	return consumer{
		querier: querier,
	}
}

// HandleMessage implements the Handler interface.
func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	ctx := context.Background()

	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}

	msg := MetricMessage{}
	err := json.Unmarshal(m.Body, &msg)
	if err != nil {
		log.Println(err)
		return nil
	}

	countParams := timescale.CountCpuUtilizationParams{
		Time:       time.Unix(msg.Time, 0),
		Type:       msg.Type,
		InstanceID: msg.InstanceId,
	}

	count, _ := h.query.CountCpuUtilization(ctx, countParams)
	if count != 0 {
		return nil
	}

	params := timescale.InsertCpuUtilizationParams{
		Time:       time.Unix(msg.Time, 0),
		InstanceID: msg.InstanceId,
		Type:       msg.Type,
		Value:      sql.NullFloat64{Float64: msg.Value, Valid: true},
	}

	_, err = h.query.InsertCpuUtilization(ctx, params)
	if err != nil {
		log.Println(err)
		return nil
	}

	return nil
}

func (c *consumer) Run() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("metrics", "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&messageHandler{c.querier})

	err = consumer.ConnectToNSQLookupd("nsqlookupd:4161")
	if err != nil {
		log.Fatal(err)
	}

	<-consumer.StopChan
	consumer.Stop()
}

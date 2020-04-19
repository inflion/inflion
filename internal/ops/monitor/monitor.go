// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package monitor

import (
	"context"
	"github.com/inflion/inflion/internal/ops/producer"
	"github.com/inflion/inflion/internal/store"
	"github.com/inflion/inflion/internal/timescale"
	"log"
	"sync"
	"time"
)

type Monitor struct {
	querier   store.Querier
	tsquerier timescale.Querier
	producer  *producer.Producer
}

func NewMonitor(querier store.Querier, tsquerier timescale.Querier, producer *producer.Producer) Monitor {
	return Monitor{
		querier:   querier,
		tsquerier: tsquerier,
		producer:  producer,
	}
}

func (m *Monitor) Run() {
	log.Println("Monitor started")
	var wg sync.WaitGroup
	wg.Add(1)

	cpu := &cpuMonitor{monitor: m}
	sg := &sgMonitor{monitor: m}

	ctx := context.Background()
	ticker := time.NewTicker(time.Minute * 6)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				cpu.run(ctx)
				sg.run(ctx)
			}
		}
	}()

	wg.Wait()
	log.Println("Monitor finished")
}

func (m *Monitor) ProduceEvent(event interface{}) {
	err := m.producer.Produce(event)
	if err != nil {
		log.Println(err)
	}
}

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
	"encoding/json"
	inflionEvent "github.com/inflion/inflion/internal/ops/event"
	"log"
)

type cpuMonitor struct {
	monitor *Monitor
}

func (m *cpuMonitor) run(ctx context.Context) {
	log.Println("cpu Monitor started")

	instance, err := m.monitor.querier.SelectInstance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range instance {
		average, _ := m.monitor.tsquerier.Average(ctx, i.InstanceID)

		// TODO thresholdを設定から取得する
		threshold := 0.01

		if average.Avg.(float64) > threshold {
			bytes, err := json.Marshal(map[string]interface{}{
				"Type":       CPUUtilization.String(),
				"Message":    "average cpu utilization is high",
				"InstanceId": i.InstanceID,
				"Value":      average.Avg.(float64),
			})
			if err != nil {
				log.Print(err)
			}

			ie, err := inflionEvent.NewInflionEvent("TODO_FIX_ME", bytes) // FIXME get project from somewhere
			if err != nil {
				log.Print(err)
			}
			m.monitor.ProduceEvent(*ie)
		}
	}
}

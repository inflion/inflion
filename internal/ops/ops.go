// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package ops

import (
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/ops/broker"
	"github.com/inflion/inflion/internal/ops/syncer"
)

type Ops struct {
	monitor monitor.Monitor
	broker  broker.Broker
	syncer  syncer.Syncer
}

func newOps(monitor monitor.Monitor, broker broker.Broker, syncer syncer.Syncer) Ops {
	return Ops{
		monitor: monitor,
		broker:  broker,
		syncer:  syncer,
	}
}

func (o *Ops) RunMonitor() {
	o.monitor.Run()
}

func (o *Ops) RunBroker() {
	o.broker.Run()
}

func (o *Ops) RunSyncer() {
	o.syncer.Run()
}

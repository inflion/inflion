//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package ops

import (
	"github.com/google/wire"
	"github.com/inflion/inflion/internal/ops/broker"
	flowstore "github.com/inflion/inflion/internal/ops/flow/store"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/ops/producer"
	"github.com/inflion/inflion/internal/ops/rule"
	"github.com/inflion/inflion/internal/ops/syncer"
	"github.com/inflion/inflion/internal/store"
	"github.com/inflion/inflion/internal/timescale"
)

func Initialize() (Ops, error) {
	wire.Build(
		wire.NewSet(
			wire.Bind(new(store.Querier), new(*store.Queries)),
			store.New,
			store.NewDbtx,
		),
		wire.NewSet(
			wire.Bind(new(timescale.Querier), new(*timescale.Queries)),
			timescale.New,
			timescale.NewDbtx,
		),
		producer.NewProducer,
		monitor.NewMonitor,
		broker.NewNsqConsumer,
		broker.NewBroker,
		broker.NewEventProcessor,
		flowstore.NewFlowStore,
		rule.NewEventMatcher,
		rule.NewRuleStore,
		syncer.NewSyncer,
		newOps,
	)
	return Ops{}, nil
}

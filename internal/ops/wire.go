//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package ops

import (
	"github.com/google/wire"
	"github.com/inflion/inflion/internal/ops/syncer"
	"github.com/inflion/inflion/internal/store"
	"github.com/inflion/inflion/internal/timescale"
	"github.com/inflion/inflion/internal/ops/notification"
	"github.com/inflion/inflion/internal/ops/monitor"
	"github.com/inflion/inflion/internal/ops/producer"
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
		notification.NewBroker,
		syncer.NewSyncer,
		newOps,
	)
	return Ops{}, nil
}

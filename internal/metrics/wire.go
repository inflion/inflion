//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package metrics

import (
	"github.com/google/wire"
	"github.com/nsqio/go-nsq"
	"github.com/inflion/inflion/internal/store"
	"github.com/inflion/inflion/internal/timescale"
)

func Initialize() (metrics, error) {
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
		wire.NewSet(
			newCollector,
			newProducer,
			newConsumer,
			newMetrics,
		),
		wire.NewSet(
			nsq.NewConfig,
		),
	)

	return metrics{}, nil
}

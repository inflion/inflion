//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package server

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/inflion/inflion/api/graphql"
	"github.com/inflion/inflion/internal/store"
)

func initServer() (server, error) {
	wire.Build(
		wire.Bind(new(store.Querier), new(*store.Queries)),
		wire.Bind(new(store.DBTX), new(*sql.DB)),
		store.New,
		store.NewDbConnection,
		graphql.NewResolverConfig,
		newServer,
		newProjectHandler,
	)
	return server{}, nil
}

// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package syncer

import (
	"context"
	"github.com/inflion/inflion/internal/logger"

	"github.com/inflion/inflion/internal/paws"
	"github.com/inflion/inflion/internal/store"
)

var log logger.Logger

func init() {
	var err error
	log, err = logger.NewZapLogger(&logger.Configuration{Level: logger.DebugLevel})
	if err != nil {
		panic(err)
	}
}

type InstanceSyncer struct {
	updater  ServiceUpdater
	fetcher  InstanceFetcher
	linker   InstanceLinker
	resolver InstanceIdResolver
	querier  Querier
}

type InstanceFetcher interface {
	GetInstances(cond paws.FilterCondition) ([]*paws.AwsInstance, error)
}

type InstanceIdResolver interface {
	ResolveIdByInstanceId(context.Context, string) (int64, error)
}

type ServiceUpdater interface {
	UpsertService(context.Context, store.UpsertServiceParams) (store.Service, error)
}

type InstanceLinker interface {
	LinkInstanceWithService(context.Context, store.LinkInstanceWithServiceParams) (store.InstanceAtService, error)
}

type Querier interface {
	CreateSecurityGroup(ctx context.Context, arg store.CreateSecurityGroupParams) error
}

func (i *InstanceSyncer) run(ctx context.Context, params SyncParams) (err error) {
	instances, err := i.fetcher.GetInstances(paws.NewEmptyFilterCondition())
	if err != nil {
		log.Error(err)
	}

	for _, instance := range instances {
		if service, ok := instance.Tags.FindValue("Service"); ok {
			input := store.UpsertServiceParams{
				Name:      service,
				ProjectID: params.projectId,
			}

			result, err := i.updater.UpsertService(ctx, input)
			if err != nil {
				log.Error(err)
				return err
			}

			for _, sg := range instance.SecurityGroups {
				err = i.querier.CreateSecurityGroup(ctx, store.CreateSecurityGroupParams{
					SecurityGroupID:   sg.Id,
					SecurityGroupName: sg.Name,
				})
			}

			if err != nil {
				log.Error(err)
			}

			id, err := i.resolver.ResolveIdByInstanceId(ctx, instance.InstanceID)
			if err != nil {
				log.Error(err)
				return err
			}

			_, err = i.linker.LinkInstanceWithService(ctx, store.LinkInstanceWithServiceParams{
				ServiceID:  result.ID,
				InstanceID: id,
			})
			if err != nil {
				log.Error(err)
			}

		}
	}

	return nil
}

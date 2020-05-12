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
	"testing"

	"github.com/inflion/inflion/internal/paws"
	"github.com/inflion/inflion/internal/store"
)

type MockInstanceIdResolver struct {
}

func (m MockInstanceIdResolver) ResolveIdByInstanceId(_ context.Context, instanceId string) (int64, error) {
	return 1, nil
}

type MockInstanceLinker struct {
}

func (m MockInstanceLinker) LinkInstanceWithService(_ context.Context, _ store.LinkInstanceWithServiceParams) (i store.InstanceAtService, er error) {
	return i, nil
}

type MockInstanceFetcher struct {
	serviceName string
}

func (m MockInstanceFetcher) GetInstances(c paws.FilterCondition) ([]*paws.AwsInstance, error) {
	return []*paws.AwsInstance{
		{
			InstanceID:     "i-test",
			Name:           "test",
			PrivateAddress: "192.168.0.1",
			PublicAddress:  "192.168.0.1",
			Tags: paws.Tags{
				Tags: []paws.Tag{
					{
						Key:   "Service",
						Value: m.serviceName,
					},
				},
			},
			SecurityGroups: nil,
			Status:         "",
		},
	}, nil
}

type MockServiceUpdater struct {
	calledName string
}

func (m *MockServiceUpdater) UpsertService(c context.Context, params store.UpsertServiceParams) (store.Service, error) {
	m.calledName = params.Name
	return store.Service{}, nil
}

func TestRun(t *testing.T) {
	ctx := context.Background()

	fetcher := MockInstanceFetcher{serviceName: "testing"}
	updater := MockServiceUpdater{}
	linker := MockInstanceLinker{}
	resolver := MockInstanceIdResolver{}

	i := InstanceSyncer{
		fetcher:  fetcher,
		updater:  &updater,
		linker:   linker,
		resolver: resolver,
	}

	err := i.run(ctx, SyncParams{})
	if err != nil {
		t.Errorf("got unexpected: %v", err)
	}

	if updater.calledName != fetcher.serviceName {
		t.Errorf("got: %v\nwant: %v", updater.calledName, fetcher.serviceName)
	}
}

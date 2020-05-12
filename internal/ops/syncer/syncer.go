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
	"sync"
	"time"

	"github.com/inflion/inflion/internal/paws"
	"github.com/inflion/inflion/internal/store"
)

const tick = time.Minute * 1

type SyncParams struct {
	awsAccount paws.AwsAccount
	projectId  int64
}

type DataRetriever interface {
	run(ctx context.Context, params SyncParams) error
}

type Syncer struct {
	querier store.Querier
}

func NewSyncer(querier store.Querier) Syncer {
	return Syncer{
		querier: querier,
	}
}

func (s *Syncer) Run() {
	ctx := context.Background()

	log.Info("Syncer started")
	var wg sync.WaitGroup
	wg.Add(1)

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	go func() {
		for {
			<-ticker.C

			awsAccounts, err := s.querier.AllAwsAccount(ctx)
			if err != nil {
				log.Error(err)
				continue
			}

			for _, account := range awsAccounts {
				params := SyncParams{
					awsAccount: paws.AwsAccount{
						AccountId:  account.AccountID,
						RoleName:   account.RoleName,
						ExternalId: account.ExternalID,
					},
					projectId: account.ProjectID,
				}

				api, err := paws.New(params.awsAccount, "ap-northeast-1")
				if err != nil {
					log.Error(err)
					continue
				}

				dataRetrievers := []DataRetriever{
					&InstanceSyncer{updater: s.querier, fetcher: &api, linker: s.querier, resolver: s.querier, querier: s.querier},
				}

				for _, r := range dataRetrievers {
					err := r.run(ctx, params)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}()

	wg.Wait()
	log.Info("Syncer finished")
}

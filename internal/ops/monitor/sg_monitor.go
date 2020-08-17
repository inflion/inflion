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
	"github.com/inflion/inflion/internal/paws"
	"log"
)

// SecurityGroup Monitor
type sgMonitor struct {
	monitor *Monitor
}

func (s *sgMonitor) run(ctx context.Context) {
	log.Println("Security group monitor started")

	awsAccounts, err := s.monitor.querier.AllAwsAccount(ctx)
	if err != nil {
		log.Println(err)
	}

	for _, account := range awsAccounts {
		awsAccount := paws.AwsAccount{
			AccountId:  account.AccountID,
			RoleName:   account.RoleName,
			ExternalId: account.ExternalID,
		}

		awsSg := paws.NewAwsSecurityGroup(awsAccount, "ap-northeast-1")
		securityGroups, err := awsSg.GetSecurityGroups()
		if err != nil {
			log.Println(err)
			continue
		}

		for _, sg := range securityGroups {
			if sg.HasOpenPorts() {
				bytes, err := json.Marshal(map[string]interface{}{
					"Type":              OpenPortDetected.String(),
					"Message":           "open port found",
					"SecurityGroupId":   sg.Id,
					"SecurityGroupName": sg.Name,
					//"OpenPorts":       sg.GetOpenPorts().ToString(),
					"OpenPorts": "22",
				})
				if err != nil {
					log.Print(err)
				}

				ie, err := inflionEvent.NewInflionEvent("TODO_FIX_ME", bytes) // FIXME get project from somewhere
				if err != nil {
					log.Print(err)
				}

				s.monitor.ProduceEvent(*ie)
			}
		}
	}
}

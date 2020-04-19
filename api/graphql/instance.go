// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package graphql

import (
	"context"
	"errors"
	"github.com/inflion/inflion/internal/hasura"
	"github.com/inflion/inflion/internal/paws"
	"github.com/inflion/inflion/internal/store"
)

func (r *queryResolver) AwsInstances(ctx context.Context, projectID int) ([]paws.AwsInstance, error) {
	userId := ctx.Value(hasura.HasuraUserIdKey).(string)
	projectId := int64(projectID)

	count, _ := r.querier.CountProjectCollaboratorByUserId(ctx, store.CountProjectCollaboratorByUserIdParams{
		UserID:    userId,
		ProjectID: projectId,
	})
	if count == 0 {
		return nil, errors.New("permission denied")
	}

	awsAccount, err := r.querier.GetAwsAccount(ctx, projectId)
	if err != nil {
		return nil, err
	}

	api, err := paws.New(
		paws.AwsAccount{
			AccountId:  awsAccount.AccountID,
			RoleName:   awsAccount.RoleName,
			ExternalId: awsAccount.ExternalID,
		},
		"ap-northeast-1",
	)
	if err != nil {
		return nil, err
	}

	instances, err := api.GetInstances(paws.NewEmptyFilterCondition())
	if err != nil {
		return nil, err
	}

	var result []paws.AwsInstance
	for _, instance := range instances {
		_, err := r.querier.UpsertInstance(ctx, store.UpsertInstanceParams{
			InstanceID: instance.InstanceID,
			Name:       instance.Name,
			ProjectID:  projectId,
			Status:     instance.Status,
		})
		if err != nil {
			return nil, err
		}

		result = append(result, paws.AwsInstance{
			InstanceID:       instance.InstanceID,
			Name:             instance.Name,
			PrivateAddress:   instance.PrivateAddress,
			PublicAddress:    instance.PublicAddress,
			Tags:             instance.Tags,
			SecurityGroupIds: instance.SecurityGroupIds,
			Status:           instance.Status,
		})
	}

	return result, nil
}

func (r *queryResolver) AwsInstance(ctx context.Context, id int) (*paws.AwsInstance, error) {
	return &paws.AwsInstance{InstanceID: "asdfasdfasdfasda", Name: "Hoge"}, nil
}

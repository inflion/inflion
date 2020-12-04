// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package paws

import "fmt"

type AwsAccount struct {
	AccountId  string
	RoleName   string
	ExternalId string
}

func (a *AwsAccount) CreateARN() string {
	return fmt.Sprintf("arn:aws:iam::%s:role/%s", a.AccountId, a.RoleName)
}

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

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/inflion/inflion/internal/logger"
)

var log logger.Logger

func init() {
	var err error
	log, err = logger.NewZapLogger(&logger.Configuration{Level: logger.DebugLevel})
	if err != nil {
		panic(err)
	}
}

func NewCloudWatch(awsAccount AwsAccount, region string) *cloudwatch.CloudWatch {
	sess, err := session.NewSession()
	if err != nil {
		log.Error(err)
	}

	conf := CreateConfig(awsAccount, region, sess)
	return cloudwatch.New(sess, &conf)
}

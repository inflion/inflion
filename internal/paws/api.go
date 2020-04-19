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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Api struct {
	conn *ec2.EC2
}

func New(awsAccount AwsAccount, region string) (api Api, err error) {
	sess, err := session.NewSession()
	if err != nil {
		return api, err
	}
	conf := CreateConfig(awsAccount, region, sess)
	return Api{conn: ec2.New(sess, &conf)}, nil
}

func CreateConfig(awsAccount AwsAccount, region string, sess *session.Session) (conf aws.Config) {
	conf = aws.Config{Region: aws.String(region)}

	// if ARN flag is passed in, we need to be able ot assume role here
	var creds *credentials.Credentials

	if awsAccount.ExternalId != "" {
		// If externalID flag is passed, we need to include it in credentials struct
		creds = stscreds.NewCredentials(sess, awsAccount.CreateARN(), func(p *stscreds.AssumeRoleProvider) {
			p.ExternalID = &awsAccount.ExternalId
		})
	} else {
		creds = stscreds.NewCredentials(sess, awsAccount.CreateARN(), func(p *stscreds.AssumeRoleProvider) {})
	}

	conf.Credentials = creds

	return
}

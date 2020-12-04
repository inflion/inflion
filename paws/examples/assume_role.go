// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	arn := "arn:aws:iam::AWS_ID:role/assumerole"
	externalId := ""
	region := "ap-northeast-1"

	conf := createConfig(arn, externalId, region, sess)

	//s3Svc := s3.New(sess, &conf)
	//var input *s3.ListBucketsInput
	//resp, _ := s3Svc.ListBuckets(input)
	//fmt.Println(resp)

	ec2Svc := ec2.New(sess, &conf)
	result, err := ec2Svc.DescribeInstances(nil)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		for _, res := range result.Reservations {
			for _, instance := range res.Instances {
				fmt.Println(instance)
			}
		}
	}
}

func createConfig(arn string, externalID string, region string, sess *session.Session) aws.Config {
	conf := aws.Config{Region: aws.String(region)}

	if arn != "" {
		// if ARN flag is passed in, we need to be able ot assume role here
		var creds *credentials.Credentials

		if externalID != "" {
			// If externalID flag is passed, we need to include it in credentials struct
			creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {
				p.ExternalID = &externalID
			})
		} else {
			creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {})
		}

		conf.Credentials = creds
	}

	return conf
}

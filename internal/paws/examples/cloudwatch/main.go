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
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)
	input := &ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{
			aws.String("sg-0522b222a60fe45a6"),
		},
	}

	result, err := svc.DescribeSecurityGroups(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func createConfig(arn string, externalID string, region string, sess *session.Session) (conf aws.Config) {
	conf = aws.Config{Region: aws.String(region)}

	if arn == "" {
		return
	}

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

	return
}

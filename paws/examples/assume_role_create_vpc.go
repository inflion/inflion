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
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	arn := "arn:aws:iam::AWS_ID:role/inflion"
	externalId := "assumerole"
	region := "ap-northeast-1"

	conf := createConfig(arn, externalId, region, sess)

	svc := ec2.New(sess, &conf)
	input := ec2.CreateVpcInput{
		CidrBlock: aws.String("10.0.0.0/16"),
	}

	result, err := svc.CreateVpc(&input)
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

	fmt.Println(result.Vpc.CidrBlockAssociationSet)

	svc.DeleteVpc(&ec2.DeleteVpcInput{VpcId: result.Vpc.VpcId})

	fmt.Println(result)
}

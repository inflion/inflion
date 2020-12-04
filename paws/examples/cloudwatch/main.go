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
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/inflion/inflion/internal/instance/infra/providers/aws/metrics"
	"time"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	arn := "arn:aws:iam::AWS_ID:role/assumerole"
	externalId := ""
	region := "ap-northeast-1"

	conf := createConfig(arn, externalId, region, sess)
	svc := cloudwatch.New(sess, &conf)

	metricNames := []string{
		"CPUUtilization",
		"StatusCheckFailed",
	}

	ns := "AWS/EC2"

	m := metrics.Metrics{CloudWatch: svc, Namespace: ns}
	md := metrics.NewMetricsData(svc, ns, metrics.TimeWindow{
		StartTime: 60 * 10 * 10 * 10,
		EndTime:   0,
		Period:    60,
	})

	metricList := m.ListMetrics(metricNames)

	results := md.ParallelGetMetricData(metricList)

	for _, result := range results {
		fmt.Printf("%s: %s at %s\n", result.Name, result.Value, time.Unix(result.Time, 0))
	}
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

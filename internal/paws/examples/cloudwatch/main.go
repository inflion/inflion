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
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/inflion/inflion/internal/paws"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cw := cloudwatch.New(sess)

	metricNames := []string{
		"CPUUtilization",
		"StatusCheckFailed",
	}

	namespace := "AWS/EC2"

	m := paws.NewMetrics(cw, namespace)
	metricList := m.ListMetrics(metricNames)

	tw := paws.TimeWindow{
		StartTime: 60 * 60 * 3,
		EndTime:   0,
		Period:    10,
	}

	md := paws.NewMetricsData(cw, namespace, tw)
	data := md.GetMetricData(metricList)
	fmt.Println(data)
}

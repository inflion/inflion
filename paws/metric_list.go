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
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"sync"
	"time"
)

type Metrics struct {
	CloudWatch *cloudwatch.CloudWatch
	Namespace  string
}

func NewMetrics(cw *cloudwatch.CloudWatch, namespace string) Metrics {
	return Metrics{
		CloudWatch: cw,
		Namespace:  namespace,
	}
}

type Metric struct {
	Name             string
	Statistics       string
	CloudWatchMetric *cloudwatch.Metric
}

func (m Metrics) ListMetrics(metricNames []string) (metricList []*Metric) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	requests := make(chan int, len(metricNames))
	for i := 0; i < len(metricNames); i++ {
		requests <- i
	}
	close(requests)

	rate := time.Millisecond * MaxRateLimitListMetrics
	throttle := time.NewTicker(rate)
	defer throttle.Stop()

	for req := range requests {
		metricName := metricNames[req]
		wg.Add(1)
		go func(metricName string) {
			defer wg.Done()
			<-throttle.C
			resp := m.ListMetric(metricName)
			mu.Lock()
			defer mu.Unlock()
			for _, metric := range resp {
				metricList = append(metricList, metric)
			}
		}(metricName)
	}

	wg.Wait()

	return
}

func (m Metrics) ListMetric(metricName string) []*Metric {
	isFirst := ""

	nextToken := aws.String(isFirst)

	listMetricsInput := &cloudwatch.ListMetricsInput{
		Namespace:  aws.String(m.Namespace),
		MetricName: aws.String(metricName),
	}

	var metricList []*cloudwatch.Metric

	for nextToken != nil {
		if *nextToken != isFirst {
			listMetricsInput.NextToken = nextToken
		}

		resp, err := m.CloudWatch.ListMetrics(listMetricsInput)
		if err != nil {
			continue
		}

		nextToken = resp.NextToken
		for _, metric := range resp.Metrics {
			if len(metric.Dimensions) != 0 {
				metricList = append(metricList, metric)
			}
		}
	}

	return m.convert(metricList)
}

func (m Metrics) convert(from []*cloudwatch.Metric) []*Metric {
	var metricList []*Metric

	for _, item := range from {
		s := *item.MetricName
		metricList = append(metricList, &Metric{Name: s, Statistics: "Maximum", CloudWatchMetric: item})
	}

	return metricList
}

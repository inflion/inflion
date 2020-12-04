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
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"sync"
	"time"
)

const MaxRateLimitListMetrics = 25
const MaxRateLimitGetMetricData = 50
const MaxMetricDataQuery = 100

type MetricsData struct {
	cloudWatch         *cloudwatch.CloudWatch
	namespace          string
	queryId2InstanceId map[string]string
	queryId2MetricName map[string]string
	timeWindow         TimeWindow
}

type MetricData struct {
	InstanceId string
	Name       string
	Time       int64
	Value      string
}

type TimeWindow struct {
	StartTime int
	EndTime   int
	Period    int
}

func NewMetricsData(cloudWatch *cloudwatch.CloudWatch, namespace string, timeWindow TimeWindow) *MetricsData {
	return &MetricsData{
		cloudWatch:         cloudWatch,
		namespace:          namespace,
		queryId2InstanceId: map[string]string{},
		queryId2MetricName: map[string]string{},
		timeWindow:         timeWindow,
	}
}

func (m MetricsData) GetMetricData(metricList []*Metric) (metricTable []MetricData) {
	endTime := aws.Time(time.Now().Add(time.Duration(int64(m.timeWindow.EndTime)) * time.Second * -1))
	startTime := aws.Time(time.Now().Add(time.Duration(int64(m.timeWindow.StartTime)) * time.Second * -1))

	resp, err := m.cloudWatch.GetMetricData(&cloudwatch.GetMetricDataInput{
		EndTime:           endTime,
		StartTime:         startTime,
		MetricDataQueries: m.createMetricDataQueries(metricList),
	})
	if err != nil {
		panic(err)
	}

	for _, result := range resp.MetricDataResults {
		for index, _ := range result.Timestamps {
			metricTable = append(metricTable,
				MetricData{
					InstanceId: m.queryId2InstanceId[*result.Id],
					Name:       m.queryId2MetricName[*result.Id],
					Time:       (*result.Timestamps[index]).Unix(),
					Value:      fmt.Sprintf("%v", *result.Values[index]),
				},
			)
		}
	}

	return
}

func (m MetricsData) ParallelGetMetricData(metricList []*Metric) (metricTable []MetricData) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	// 1回のGetMetricDataで取得するMetricDataQuery用のデータを準備
	bulkNum := MaxMetricDataQuery / len(metricList)
	var metricParams [][]*Metric
	var metricParam []*Metric

	for idx, metric := range metricList {
		metricParam = append(metricParam, metric)
		if idx%bulkNum == bulkNum-1 || idx == len(metricList)-1 {
			metricParams = append(metricParams, metricParam)
			metricParam = []*Metric{}
		}
	}

	// rate-limit. see https://gobyexample.com/rate-limiting
	requests := make(chan int, len(metricParams))
	for i := 0; i < len(metricParams); i++ {
		requests <- i
	}
	close(requests)

	rate := time.Millisecond * MaxRateLimitGetMetricData
	throttle := time.NewTicker(rate)
	defer throttle.Stop()

	for req := range requests {
		bulkMetric := metricParams[req]

		wg.Add(1)

		go func(bulkMetric []*Metric) {
			defer wg.Done()
			<-throttle.C
			resp := m.GetMetricData(bulkMetric)
			mu.Lock()
			defer mu.Unlock()
			for _, metricRecord := range resp {
				metricTable = append(metricTable, metricRecord)
			}
		}(bulkMetric)
	}

	wg.Wait()

	return
}

func (m MetricsData) createMetricDataQueries(metricList []*Metric) []*cloudwatch.MetricDataQuery {
	var metricDataQueries []*cloudwatch.MetricDataQuery

	for i, metric := range metricList {
		metricDataId := fmt.Sprintf("m%d", i)

		m.queryId2InstanceId[metricDataId] = *metric.CloudWatchMetric.Dimensions[0].Value
		m.queryId2MetricName[metricDataId] = metric.Name

		metricDataQueries = append(
			metricDataQueries,
			m.createMetricDataQuery(metricDataId, metric),
		)
	}

	return metricDataQueries
}

func (m MetricsData) createMetricDataQuery(metricDataId string, metric *Metric) *cloudwatch.MetricDataQuery {
	return &cloudwatch.MetricDataQuery{
		Id: aws.String(metricDataId),
		MetricStat: &cloudwatch.MetricStat{
			Metric: &cloudwatch.Metric{
				Namespace:  metric.CloudWatchMetric.Namespace,
				Dimensions: metric.CloudWatchMetric.Dimensions,
				MetricName: &metric.Name,
			},
			Period: aws.Int64(int64(m.timeWindow.Period)),
			Stat:   &metric.Statistics,
		},
	}
}

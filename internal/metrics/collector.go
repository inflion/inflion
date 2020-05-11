// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package metrics

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/inflion/inflion/internal/paws"
	"github.com/inflion/inflion/internal/store"
	"log"
	"strconv"
	"sync"
	"time"
)

type collector struct {
	producer *producer
	querier  store.Querier
	sessions map[string]*cloudwatch.CloudWatch
}

type MetricMessage struct {
	InstanceId string  `json:"instance_id"`
	Type       string  `json:"type"`
	Value      float64 `json:"value"`
	Time       int64   `json:"time"`
}

var metricNames []string
var namespace string
var region string

func init() {
	metricNames = []string{
		"CPUUtilization",
		"StatusCheckFailed",
		"NetworkIn",
		"NetworkOut",
		"NetworkPacketsIn",
		"NetworkPacketsOut",
		"DiskReadOps",
		"DiskReadBytes",
		"DiskWriteOps",
		"DiskWriteBytes",
	}
	namespace = "AWS/EC2"
	region = "ap-northeast-1"
}

func newCollector(querier store.Querier, producer *producer) collector {
	return collector{
		querier:  querier,
		producer: producer,
		sessions: make(map[string]*cloudwatch.CloudWatch),
	}
}

func (c *collector) Run() {
	log.Println("Collector started")

	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		ticker := time.NewTicker(time.Second * 60)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				awsAccounts, err := c.querier.AllAwsAccount(ctx)
				if err != nil {
					log.Println(err)
				}

				for _, account := range awsAccounts {
					cw := c.getSession(account)
					metricList := c.getMetricList(cw)
					metrics := c.getMetrics(cw, metricList)
					c.produceMetrics(metrics)
				}
			}
		}
	}()

	wg.Wait()
}

func (c *collector) getSession(account store.AllAwsAccountRow) *cloudwatch.CloudWatch {
	found, ok := c.sessions[account.AccountID]
	if ok {
		return found
	} else {
		cw := paws.NewCloudWatch(
			paws.AwsAccount{
				AccountId:  account.AccountID,
				RoleName:   account.RoleName,
				ExternalId: account.ExternalID,
			},
			region,
		)
		c.sessions[account.AccountID] = cw
	}

	cw, ok := c.sessions[account.AccountID]
	if !ok {
		log.Fatal("assertion failure")
	}
	return cw
}

func (c *collector) getMetricList(cw *cloudwatch.CloudWatch) (metricList []*paws.Metric) {
	m := paws.NewMetrics(cw, namespace)
	return m.ListMetrics(metricNames)
}

func (c *collector) getMetrics(cw *cloudwatch.CloudWatch, metricList []*paws.Metric) []paws.MetricData {
	md := paws.NewMetricsData(cw, namespace, paws.TimeWindow{
		StartTime: 60 * 60 * 3,
		EndTime:   0,
		Period:    10,
	})

	return md.ParallelGetMetricData(metricList)
}

func (c *collector) produceMetrics(metricTable []paws.MetricData) {
	for _, result := range metricTable {
		v, err := strconv.ParseFloat(result.Value, 64)

		m := MetricMessage{
			Time:       result.Time,
			Type:       result.Name,
			InstanceId: result.InstanceId,
			Value:      v,
		}

		bytes, err := json.Marshal(m)
		if err != nil {
			log.Println(err)
			continue
		}

		err = c.producer.produce(bytes)
		if err != nil {
			log.Println(err)
		}
	}
}

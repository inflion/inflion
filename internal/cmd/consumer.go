// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package cmd

import (
	"github.com/inflion/inflion/internal/metrics"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(metricsConsumerCommand())
}

func metricsConsumerCommand() *cobra.Command {
	lpc := &cobra.Command{
		Use:   "consumer <subcommand>",
		Short: "consumer related command",
	}
	lpc.AddCommand(metricsConsumerStartCommand())

	return lpc
}

func metricsConsumerStartCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "start the consumer daemon",
		Run:   startMetricsConsumer,
	}

	return &cmd
}

func startMetricsConsumer(cmd *cobra.Command, args []string) {
	m, err := metrics.Initialize()
	if err != nil {
		log.Fatal("initialization failed")
	}
	m.Consumer.Run()
}

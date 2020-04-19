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
	rootCmd.AddCommand(metricsCollectorCommand())
}

func metricsCollectorCommand() *cobra.Command {
	lpc := &cobra.Command{
		Use:   "collector <subcommand>",
		Short: "collector is collect metrics",
	}
	lpc.AddCommand(metricsCollectorStartCommand())

	return lpc
}

func metricsCollectorStartCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "start the metrics collector daemon",
		Run:   startMetricsCollector,
	}

	return &cmd
}

func startMetricsCollector(cmd *cobra.Command, args []string) {
	m, err := metrics.Initialize()
	if err != nil {
		log.Fatal("initialization failed")
	}
	m.Collector.Run()
}

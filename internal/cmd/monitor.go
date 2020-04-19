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
	"github.com/inflion/inflion/internal/ops"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(monitorCommand())
}

func monitorCommand() *cobra.Command {
	lpc := &cobra.Command{
		Use:   "monitor <subcommand>",
		Short: "monitor related command",
	}
	lpc.AddCommand(monitorStartCommand())

	return lpc
}

func monitorStartCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "start the monitor",
		Run:   startMonitorCollector,
	}

	return &cmd
}

func startMonitorCollector(cmd *cobra.Command, args []string) {
	o, err := ops.Initialize()
	if err != nil {
		log.Fatal("initialization failed")
	}
	o.RunMonitor()
}

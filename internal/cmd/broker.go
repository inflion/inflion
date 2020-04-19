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
	rootCmd.AddCommand(brokerCommand())
}

func brokerCommand() *cobra.Command {
	lpc := &cobra.Command{
		Use:   "broker <subcommand>",
		Short: "broker related command",
	}
	lpc.AddCommand(brokerStartCommand())

	return lpc
}

func brokerStartCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "start the monitor",
		Run:   startBroker,
	}

	return &cmd
}

func startBroker(cmd *cobra.Command, args []string) {
	o, err := ops.Initialize()
	if err != nil {
		log.Fatal("initialization failed")
	}
	o.RunBroker()
}

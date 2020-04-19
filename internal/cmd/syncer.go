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
	rootCmd.AddCommand(syncerCommand())
}

func syncerCommand() *cobra.Command {
	lpc := &cobra.Command{
		Use:   "syncer <subcommand>",
		Short: "syncer related command",
	}
	lpc.AddCommand(syncerStartCommand())

	return lpc
}

func syncerStartCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "start the monitor",
		Run:   startSyncer,
	}

	return &cmd
}

func startSyncer(cmd *cobra.Command, args []string) {
	o, err := ops.Initialize()
	if err != nil {
		log.Fatal("initialization failed")
	}
	o.RunSyncer()
}

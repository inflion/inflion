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
	"github.com/inflion/inflion/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCommand())
}

func serverCommand() *cobra.Command {
	lpc := &cobra.Command{
		Use:   "server <subcommand>",
		Short: "server",
	}
	lpc.AddCommand(serverStartCommand())

	return lpc
}

func serverStartCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "start api server",
		Run:   startServer,
	}

	return &cmd
}

func startServer(cmd *cobra.Command, args []string) {
	server.Run()
}

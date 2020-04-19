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
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:        "inflion",
		Short:      "inflion",
		SuggestFor: []string{"inflion"},
	}
)

func Main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	return
}

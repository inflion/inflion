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
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:        "lionctl",
		Short:      "lionctl",
		SuggestFor: []string{"lionctl"},
	}
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".inflion")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("endpoint", "e", "", "an address of a endpoint")
	err := viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringP("project", "p", "", "an name of a project")
	err = viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(flowCommand())
	rootCmd.AddCommand(ruleCommand())
	rootCmd.AddCommand(jobCommand())
}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	return
}

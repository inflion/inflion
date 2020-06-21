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
	"github.com/inflion/inflion/lionctl/clientv1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

var (
	id   string
	body string
	file string
)

func flowCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flow <subcommand>",
		Short: "flow related commands",
	}

	cmd.AddCommand(flowListCommand())
	cmd.AddCommand(flowRunCommand())
	cmd.AddCommand(flowCreateCommand())
	cmd.AddCommand(flowGetCommand())
	cmd.AddCommand(flowUpdateCommand())
	cmd.AddCommand(flowRemoveCommand())

	return cmd
}

func project() string {
	return viper.GetString("project")
}

func endpoint() string {
	return viper.GetString("endpoint")
}

func flowListCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "list flows",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewFlowClient(project(), endpoint())
			flows, err := c.List()
			if err != nil {
				log.Fatal(err)
			}

			for _, v := range flows {
				fmt.Println(v.Id)
			}

			os.Exit(0)
		},
	}

	return &cmd
}

func flowRunCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "run",
		Short: "run flow",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewFlowClient(project(), endpoint())
			output, err := c.Run(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("A flow was ran successfully: " + output)
			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id of a flow")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

func flowCreateCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "create",
		Short: "Create a new flow",
		Run: func(cmd *cobra.Command, args []string) {
			flowjson, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}

			c := clientv1.NewFlowClient(project(), endpoint())
			flow, err := c.Create(string(flowjson))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("A flow was created successfully: " + flow.Id)
			os.Exit(0)
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "path of a flow file(required)")
	err := cmd.MarkFlagRequired("file")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

func flowGetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "Get a flow",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewFlowClient(project(), endpoint())
			flow, err := c.Get(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(flow.Body)

			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id of a flow")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

func flowUpdateCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "update",
		Short: "Update a flow",
		Run: func(cmd *cobra.Command, args []string) {
			flowjson, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}

			c := clientv1.NewFlowClient(project(), endpoint())
			id, err := c.Update(id, string(flowjson))
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}

			fmt.Println("Flow was updated successfully: " + id)
			os.Exit(0)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "an id of a flow")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "path of a flow file(required)")
	err = cmd.MarkFlagRequired("file")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

func flowRemoveCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "remove",
		Short: "Remove a flow",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewFlowClient(project(), endpoint())

			id, err := c.Remove(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Flow was removed successfully: " + id)
			os.Exit(0)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "id of a flow")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

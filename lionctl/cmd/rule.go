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
	"io/ioutil"
	"log"
	"os"
)

func ruleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rule <subcommand>",
		Short: "rule related commands",
	}

	cmd.AddCommand(ruleCreateCommand())
	cmd.AddCommand(ruleGetCommand())
	cmd.AddCommand(ruleUpdateCommand())
	cmd.AddCommand(ruleRemoveCommand())

	return cmd
}

func ruleCreateCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "create",
		Short: "Create a new rule",
		Run: func(cmd *cobra.Command, args []string) {
			rulejson, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}

			c := clientv1.NewRuleClient(project(), endpoint())
			rule, err := c.Create(string(rulejson))
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("A rule was created successfully: " + rule.Id)
			os.Exit(0)
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "path of a rule file(required)")
	err := cmd.MarkFlagRequired("file")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

func ruleGetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "Get a rule",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewRuleClient(project(), endpoint())
			rule, err := c.Get(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(rule.Body)

			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id of a rule")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

func ruleUpdateCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "update",
		Short: "Update a rule",
		Run: func(cmd *cobra.Command, args []string) {
			rulejson, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}

			c := clientv1.NewRuleClient(project(), endpoint())
			id, err := c.Update(id, string(rulejson))
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}

			fmt.Println("Rule was updated successfully: " + id)
			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "an id of a rule")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "path of a rule file(required)")
	err = cmd.MarkFlagRequired("file")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

func ruleRemoveCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "remove",
		Short: "Remove a rule",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewRuleClient(project(), endpoint())

			id, err := c.Remove(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Rule was removed successfully: " + id)
			os.Exit(0)
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "id of a rule")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

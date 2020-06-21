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
	"log"
	"os"
	"strconv"
)

func jobCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "job <subcommand>",
		Short: "job related commands",
	}

	cmd.AddCommand(jobListCommand())
	cmd.AddCommand(jobCreateCommand())
	cmd.AddCommand(jobRemoveCommand())

	return cmd
}

func jobListCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "list jobs",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewJobClient(endpoint())

			project, _ := cmd.Flags().GetString("project")

			r, err := c.List(project)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Job ID, Schedule, Flow")
			for _, j := range r {
				fmt.Printf("%d, %s, %s\n", j.Id, j.Schedule, j.FlowId)
			}

			os.Exit(0)
		},
	}

	return &cmd
}

func jobCreateCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "create",
		Short: "Create a new job",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewJobClient(endpoint())

			id, _ := cmd.Flags().GetString("id")
			project, _ := cmd.Flags().GetString("project")
			flowId, _ := cmd.Flags().GetString("flow_id")
			schedule, _ := cmd.Flags().GetString("schedule")

			intId, err := strconv.ParseInt(id, 10, 32)
			if err != nil {
				log.Println(err)
				return
			}
			err = c.Create(clientv1.Job{
				Id:       int32(intId),
				Project:  project,
				FlowId:   flowId,
				Schedule: schedule,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("A job was created successfully")
			os.Exit(0)
		},
	}

	cmd.Flags().StringP("id", "i", "", "an id of a job")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}
	cmd.Flags().StringP("flow_id", "f", "", "an id of a flow")
	err = cmd.MarkFlagRequired("flow_id")
	if err != nil {
		log.Fatal(err)
	}
	cmd.Flags().StringP("schedule", "s", "", "cron format schedule")
	err = cmd.MarkFlagRequired("schedule")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

func jobRemoveCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "remove",
		Short: "Remove a job",
		Run: func(cmd *cobra.Command, args []string) {
			c := clientv1.NewJobClient(endpoint())

			id, _ := cmd.Flags().GetString("id")
			project, _ := cmd.Flags().GetString("project")

			intId, err := strconv.ParseInt(id, 10, 32)
			if err != nil {
				log.Println(err)
				return
			}
			err = c.Remove(clientv1.Job{
				Id:      int32(intId),
				Project: project,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("A job was removed successfully")
			os.Exit(0)
		},
	}

	cmd.Flags().StringP("id", "i", "", "an id of a job")
	err := cmd.MarkFlagRequired("id")
	if err != nil {
		log.Fatal(err)
	}

	return &cmd
}

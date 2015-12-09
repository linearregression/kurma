// Copyright 2015 Apcera Inc. All rights reserved.

package commands

import (
	"fmt"
	"os"

	"github.com/apcera/kurma/client/cli"
	"github.com/spf13/cobra"
)

var (
	StopCmd = &cobra.Command{
		Use:   "stop UUID",
		Short: "Stop a running container",
		Run:   cmdStop,
	}
)

func init() {
	cli.RootCmd.AddCommand(StopCmd)
}

func cmdStop(cmd *cobra.Command, args []string) {
	if len(args) == 0 || len(args) > 1 {
		fmt.Printf("Invalid command options specified.\n")
		cmd.Help()
		return
	}

	if err := cli.GetClient().DestroyContainer(args[0]); err != nil {
		fmt.Printf("Failed to stop the container: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Destroyed container %s\n", args[0])
}

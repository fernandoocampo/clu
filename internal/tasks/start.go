package tasks

import (
	"log"

	"github.com/spf13/cobra"
)

var startWork *cobra.Command

const (
	startCmdName   = "start"
	startShortDesc = "start creates a folder in your task workspace related to your new assignment"
)

func makeStartCommand() *cobra.Command {
	newTaskCommand := cobra.Command{
		Use:   startCmdName,
		Short: startShortDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				log.Fatalf("running help command: %s", err)
			}
		},
	}

	return &newTaskCommand
}

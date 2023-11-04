package tasks

import (
	"log"

	"github.com/spf13/cobra"
)

var ticket *cobra.Command

const (
	cmdName   = "task"
	shortDesc = "tasks deals with daily tasks and tickets"
)

func MakeCommand() *cobra.Command {
	newTaskCommand := &cobra.Command{
		Use:   cmdName,
		Short: shortDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				log.Fatalf("running help command: %s", err)
			}
		},
	}

	addSubcommands(newTaskCommand)

	return newTaskCommand
}

func addSubcommands(parent *cobra.Command) {
	parent.AddCommand(makeStartCommand())
}

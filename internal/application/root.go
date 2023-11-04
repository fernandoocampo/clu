package application

import (
	"fmt"
	"log"
	"os"

	"github.com/fernandoocampo/clu/internal/tasks"
	"github.com/spf13/cobra"
)

const (
	appName   = "clu"
	shortDesc = "clu is a developer assistant"
	longDesc  = "clu is a developer assistant to help you with daily tasks"
)

var (
	version = "unknown"
)

var (
	rootCommand *cobra.Command
	errorLogger *log.Logger
)

func init() {
	errorLogger = log.New(os.Stderr, "", 0)
	rootCommand = makeRootCommand()
}

func Run() error {
	err := rootCommand.Execute()
	if err != nil {
		errorLogger.Println("executing clu command: ", err)

		return fmt.Errorf("unable to execute root command: %w", err)
	}

	return nil
}

func makeRootCommand() *cobra.Command {
	newRootCommand := &cobra.Command{
		Use:     appName,
		Short:   shortDesc,
		Version: version,
		Long:    longDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				log.Fatalf("running help command: %s", err)
			}
		},
	}

	addSubcommands(newRootCommand)

	return newRootCommand
}

func addSubcommands(rootCommand *cobra.Command) {
	rootCommand.AddCommand(tasks.MakeCommand())
}

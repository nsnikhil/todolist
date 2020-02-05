package cmd

import (
	"github.com/spf13/cobra"
)

const (
	appName        = "todolist"
	appDescription = "A todolist management application"
	appVersion     = "0.1-alpha"

	serveCommandName        = "serve"
	serveCommandDescription = "start the application"

	migrateCommandName        = "migrate"
	migrateCommandDescription = "run migrations"

	rollbackCommandName        = "rollback"
	rollbackCommandDescription = "rollback migrations"

	configCommandName        = "config"
	configCommandDescription = "print currently loaded configuration"
)

func newCommand(name string, description string, run func()) *cobra.Command {
	return &cobra.Command{
		Use:   name,
		Short: description,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
}

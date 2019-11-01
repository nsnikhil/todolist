package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"todolist/applogger"
	"todolist/constants"
	"todolist/store"
)

var migrateCmd = newCommand(constants.MigrateCommandName, constants.MigrateCommandDescription, runMigrations)

func runMigrations() {
	if err := store.Migrate(); err != nil {
		if err == migrate.ErrNoChange {
			applogger.Infof(constants.SchemaUpToDate, "[runMigrations] [Migrate]")
			return
		}
		applogger.Errorf(constants.ErrorMigrationFailed, "[runMigrations] [Migrate]", err)
		return
	}
	applogger.Infof(constants.SuccessfulMigrations, "[runMigrations] [Migrate]")
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

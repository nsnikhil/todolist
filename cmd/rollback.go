package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"todolist/applogger"
	"todolist/constants"
	"todolist/store"
)

var rollbackCmd = newCommand(constants.RollbackCommandName, constants.RollbackCommandDescription, rollBackMigrations)

func rollBackMigrations() {
	if err := store.Rollback(); err != nil {
		if err == migrate.ErrNoChange {
			applogger.Infof(constants.SchemaUpToDate, "[rollBackMigrations] [Rollback]")
			return
		}
		applogger.Errorf(constants.ErrorRollbackFailed, "[rollBackMigrations] [Rollback]", err)
	}
	applogger.Infof(constants.SuccessfulRollback, "[rollBackMigrations] [Rollback]")
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}

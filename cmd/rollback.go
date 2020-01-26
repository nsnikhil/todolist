package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"todolist/store"
	"todolist/util"
)

var rollbackCmd = newCommand(rollbackCommandName, rollbackCommandDescription, rollBackMigrations)

func rollBackMigrations() {
	if err := store.Rollback(); err != nil {
		if err == migrate.ErrNoChange {
			util.DebugLog("[rollBackMigrations] [ErrNoChange]")
			return
		}
		util.LogError("[rollBackMigrations] [Rollback]", err)
	}
	util.DebugLog("[rollBackMigrations] [Rollback] [Success]")
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}

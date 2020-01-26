package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"todolist/store"
	"todolist/util"
)

var migrateCmd = newCommand(migrateCommandName, migrateCommandDescription, runMigrations)

func runMigrations() {
	if err := store.Migrate(); err != nil {
		if err == migrate.ErrNoChange {
			util.DebugLog("[runMigrations] [ErrNoChange]")
			return
		}
		util.LogError("[runMigrations] [Migrate]", err)
		return
	}
	util.DebugLog("[runMigrations] [Migrate] [Success]")
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

package store

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"path/filepath"
	"strings"
	"todolist/config"
	"todolist/constants"
	"todolist/util"
)

const (
	migrationPath = "./store/migrations"
	rollBackStep  = -1
	cutSet        = "file://"
)

func Migrate() error {
	newMigrate, err := newMigrate()
	if err != nil {
		return util.LogAndGetError("[Migrate] [newMigrate]", err)
	}
	return newMigrate.Up()
}

func Rollback() error {
	newMigrate, err := newMigrate()
	if err != nil {
		return util.LogAndGetError("[Migrate] [newMigrate]", err)
	}
	return newMigrate.Steps(rollBackStep)
}

func newMigrate() (*migrate.Migrate, error) {
	if err := config.Load(); err != nil {
		return nil, util.LogAndGetError("[newMigrate] [config.Load]", err)
	}

	dbHandler := NewDBHandler(config.GetDatabaseConfig())

	db, err := dbHandler.GetDB()

	if err != nil {
		return nil, util.LogAndGetError("[newMigrate] [GetDB]", err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, util.LogAndGetError("[newMigrate] [postgres.WithInstance]", err)
	}

	sourcePath, err := getSourcePath(migrationPath)
	if err != nil {
		return nil, util.LogAndGetError("[newMigrate] [getSourcePath]", err)
	}

	util.DebugLog("[newMigrate] [getSourcePath]", sourcePath)
	return migrate.NewWithDatabaseInstance(sourcePath, "postgres", driver)
}

func getSourcePath(directory string) (string, error) {
	directory = strings.TrimLeft(directory, cutSet)
	absPath, err := filepath.Abs(directory)
	if err != nil {
		return constants.EmptyString, err
	}
	return fmt.Sprintf("%s%s", cutSet, absPath), nil
}

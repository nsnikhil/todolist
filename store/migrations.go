package store

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"path/filepath"
	"strings"
	"todolist/applogger"
	"todolist/config"
	"todolist/constants"
)

func Migrate() error {
	newMigrate, err := newMigrate()
	if err != nil {
		applogger.Errorf(constants.ErrorFailedToGetMigrate, "[Migrate] [newMigrate]", err)
		return err
	}
	return newMigrate.Up()
}

func Rollback() error {
	newMigrate, err := newMigrate()
	if err != nil {
		applogger.Errorf(constants.ErrorFailedToGetMigrate, "[Migrate] [newMigrate]", err)
		return err
	}
	return newMigrate.Steps(constants.RollBackStep)
}

func newMigrate() (*migrate.Migrate, error) {
	if err := config.Load(); err != nil {
		applogger.Errorf(constants.ErrorFailedToLoadConfig, "[runMigrations] [Load]", err)
		return nil, err
	}

	dbHandler := NewDBHandle(config.GetDatabaseConfig())

	db, err := dbHandler.GetDB()

	if err != nil {
		applogger.Errorf(constants.ErrorDatabaseFailedToLoad, fmt.Sprint("[runMigrations] [GetDB]"), config.GetDatabaseConfig().String(), err)
		return nil, err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		applogger.Errorf(constants.ErrorFailedToGetDatabaseDriver, "[runMigrations] [WithInstance]", err)
		return nil, err
	}

	sourcePath, err := getSourcePath(constants.MigrationPath)
	if err != nil {
		return nil, err
	}
	applogger.Infof(constants.SchemaPath, "[runMigrations] [getSourcePath]", sourcePath)

	return migrate.NewWithDatabaseInstance(sourcePath, "postgres", driver)
}

func getSourcePath(directory string) (string, error) {
	directory = strings.TrimLeft(directory, constants.CutSet)
	absPath, err := filepath.Abs(directory)
	if err != nil {
		return constants.EmptyString, err
	}
	return fmt.Sprintf("%s%s", constants.CutSet, absPath), nil
}

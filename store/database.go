package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"todolist/apperror"
	"todolist/applogger"
	"todolist/config"
	"todolist/constants"
)

type DBHandleInterface interface {
	GetDB() (*sqlx.DB, error)
}

type DBHandle struct {
	config.DatabaseConfig
}

func NewDBHandle(config config.DatabaseConfig) DBHandle {
	applogger.Infof(constants.DBHandleNew, fmt.Sprint("[DBHandle] [NewDBHandle]"), config.String())
	return DBHandle{DatabaseConfig: config}
}

func (dbh DBHandle) GetDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(constants.DatabaseDriverName, dbh.String())
	if err != nil {
		applogger.Errorf(constants.ErrorDatabaseFailedToLoad, fmt.Sprint("[DBHandle] [GetDB]"), dbh.String(), err)
		return nil, apperror.NewDatabaseLoadError(constants.ErrorDatabaseFailedToLoad, fmt.Sprint("[DBHandle] [GetDB]"), dbh.String(), err.Error())
	}

	if err = db.Ping(); err != nil {
		applogger.Errorf(constants.ErrorDatabasePingFailed, fmt.Sprint("[DBHandle] [GetDB]"), err)
		return nil, apperror.NewDatabasePingError(constants.ErrorDatabasePingFailed, fmt.Sprint("[DBHandle] [GetDB]"), err.Error())
	}

	db.SetMaxOpenConns(dbh.GetMaxPoolSize())
	applogger.Infof(constants.SuccessfulConnectionToDatabase, fmt.Sprint("[DBHandle] [GetDB]"), dbh.String(), dbh.GetMaxPoolSize())
	return db, nil
}

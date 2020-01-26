package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"todolist/config"
	"todolist/util"
)

const databaseDriverName  = "postgres"

type DBHandler interface {
	GetDB() (*sqlx.DB, error)
}

type DefaultDBHandler struct {
	config.DatabaseConfig
}

func NewDBHandler(config config.DatabaseConfig) DBHandler {
	util.DebugLog("[DefaultDBHandler] [NewDBHandler]", config)
	return DefaultDBHandler{DatabaseConfig: config}
}

func (dbh DefaultDBHandler) GetDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(databaseDriverName, dbh.Source())

	if err != nil {
		return nil, util.LogAndGetError("[DefaultDBHandler] [GetDB]", errors.Wrap(err, fmt.Sprintf("failed to server db for %s", dbh.Source())))
	}

	if err = db.Ping(); err != nil {
		return nil, util.LogAndGetError("[DefaultDBHandler] [Ping]", err)
	}

	db.SetMaxOpenConns(dbh.GetMaxPoolSize())
	util.DebugLog("[DefaultDBHandler] [GetDB] [Success]")
	return db, nil
}

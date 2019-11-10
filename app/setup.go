package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	"todolist/applogger"
	"todolist/config"
	"todolist/constants"
	"todolist/service"
	"todolist/store"
)

func SetUpDependencies() Dependencies {
	applogger.Infof(constants.SetupDependencies, "[SetUpDependencies]")
	return Dependencies{
		Service: setupService(),
	}
}

func setupDBConnection() *sqlx.DB {

	dbConfig := config.GetDatabaseConfig()

	dbHandler := store.NewDBHandle(dbConfig)

	db, err := getDBWithRetry(dbHandler, dbConfig.GetRetryCount())

	if err != nil {
		applogger.Errorf(constants.ErrorDatabaseFailedToLoad, fmt.Sprint("[setupDBConnection] [GetDB]"), config.GetDatabaseConfig().String(), err)
		panic(err)
	}

	applogger.Infof(constants.SetupDB, "[setupDBConnection]", db)

	return db
}

func getDBWithRetry(dbHandler store.DBHandleInterface, retryCount int) (*sqlx.DB, error) {
	applogger.Infof(constants.GetDBWithRetry, "[getDBWithRetry]", retryCount)
	for i := retryCount; i > 0; i-- {
		db, err := dbHandler.GetDB()
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}
	return dbHandler.GetDB()
}

func setupStore() store.Store {
	applogger.Infof(constants.SetupStore, "[setupStore]")
	return store.NewStore(setupDBConnection())
}

func setupService() service.Service {
	todoListService := service.NewTodoListService(setupStore().GetTodoListStore())
	applogger.Infof(constants.SetupService, "[setupService]")
	return service.NewService(todoListService)
}

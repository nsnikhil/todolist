package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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

	dbHandler := store.NewDBHandle(config.GetDatabaseConfig())

	db, err := dbHandler.GetDB()

	if err != nil {
		applogger.Errorf(constants.ErrorDatabaseFailedToLoad, fmt.Sprint("[setupDBConnection] [GetDB]"), config.GetDatabaseConfig().String(), err)
		panic(err)
	}

	applogger.Infof(constants.SetupDB, "[setupDBConnection]", db)

	return db
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

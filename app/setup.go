package app

import (
	"github.com/jmoiron/sqlx"
	"time"
	"todolist/config"
	"todolist/domain"
	"todolist/service"
	"todolist/store"
	"todolist/util"
)

func SetUpDependencies() Dependencies {
	util.DebugLog("[SetUpDependencies]")
	return Dependencies{
		TaskFactory: setUpFactory(),
		Service:     setupService(),
	}
}

func setUpFactory() domain.TaskFactory {
	return domain.NewTaskFactory()
}

func setupDBConnection() *sqlx.DB {
	dbConfig := config.GetDatabaseConfig()
	dbHandler := store.NewDBHandler(dbConfig)

	db, err := getDBWithRetry(dbHandler, dbConfig.GetRetryCount())
	if err != nil {
		panic(util.LogAndGetError("[setupDBConnection] [getDBWithRetry]", err))
	}

	util.DebugLog("[setupDBConnection] [Success]")
	return db
}

func getDBWithRetry(dbHandler store.DBHandler, retryCount int) (*sqlx.DB, error) {
	util.DebugLog("[getDBWithRetry] : ", retryCount)
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
	util.DebugLog("[setupStore]")
	return store.NewStore(setupDBConnection())
}

func setupService() service.Service {
	util.DebugLog("[setupService]")
	return service.NewService(service.NewTaskService(setupStore().GetTodoListStore()))
}

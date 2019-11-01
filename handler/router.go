package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"todolist/app"
	"todolist/applogger"
	"todolist/constants"
)

func NewRouter(dependencies app.Dependencies) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc(constants.PathPing, pingHandler()).Methods(http.MethodGet)

	todolistRouter := router.PathPrefix(constants.PathTodoList).Subrouter()
	todolistRouter.HandleFunc(constants.PathAdd, addHandler(dependencies.GetTodoListService())).Methods(http.MethodPost)
	todolistRouter.HandleFunc(constants.PathRemove, removeHandler(dependencies.GetTodoListService())).Methods(http.MethodDelete)
	todolistRouter.HandleFunc(constants.PathUpdate, updateHandler(dependencies.GetTodoListService())).Methods(http.MethodPost)
	todolistRouter.HandleFunc(constants.PathGetTodoList, getTodoListHandler(dependencies.GetTodoListService())).Methods(http.MethodGet)
	todolistRouter.HandleFunc(constants.PathGetTask, getTaskHandler(dependencies.GetTodoListService())).Methods(http.MethodGet)

	applogger.Infof(constants.NewRouterCreated, "[NewRouter]", router)

	return http.HandlerFunc(router.ServeHTTP)
}

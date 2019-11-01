package handler

import (
	"net/http"
	"todolist/applogger"
	"todolist/constants"
	"todolist/contract"
	"todolist/service"
)

func getTodoListHandler(service service.TodoListServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var response contract.TodoListResponse

		todoList, err := service.GetTodoList()
		if err != nil {
			applogger.Errorf(constants.ErrorFailedToGetAllTasks, "[getTodoListHandler] [GetTodoList]", err)
			writeErrorResponse(w, http.StatusInternalServerError, err, "getTodoListHandler")
			return
		}

		response = contract.NewTodoListResponse(todoList, constants.EmptyString, true)

		applogger.Infof(constants.SuccessfulGetAllTask, "[getTodoListHandler]", todoList)

		writeResponse(w, http.StatusOK, response, "getTodoListHandler")

	}
}

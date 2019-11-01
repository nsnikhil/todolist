package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"todolist/applogger"
	"todolist/constants"
	"todolist/contract"
	"todolist/service"
)

func getTaskHandler(service service.TodoListServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var response contract.TodoListResponse

		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			applogger.Errorf(constants.ErrorFailedToReadRequestBody, "[getTaskHandler] [ReadAll]", err)
			writeErrorResponse(w, http.StatusBadRequest, err, "getTaskHandler")
			return
		}

		var getTaskRequest contract.TaskIDRequest

		if err = json.Unmarshal(bytes, &getTaskRequest); err != nil {
			applogger.Errorf(constants.ErrorFailedToUnMarshalRequestBody, "[getTaskHandler] [Unmarshal]", bytes, err)
			writeErrorResponse(w, http.StatusBadRequest, err, "getTaskHandler")
			return
		}

		if !getTaskRequest.Validate() {
			applogger.Errorf(constants.ErrorRequestBodyValidationFailed, "[getTaskHandler] [Validate]")
			writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf(constants.ErrorTaskValidationFailed, "[getTaskHandler] [Validate]", getTaskRequest.TaskID), "getTaskHandler")
			return
		}

		task, err := service.GetTask(getTaskRequest.TaskID)
		if err != nil {
			applogger.Errorf(constants.ErrorTaskNotFound, "[getTaskHandler] [GetTask]", getTaskRequest.TaskID, err)
			writeErrorResponse(w, http.StatusInternalServerError, err, "getTaskHandler")
			return
		}

		response = contract.NewTodoListResponse(task, constants.EmptyString, true)

		applogger.Infof(constants.SuccessfulGetTask, "[getTodoListHandler]", task)

		writeResponse(w, http.StatusOK, response, "getTodoListHandler")

	}
}

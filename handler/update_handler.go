package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"todolist/applogger"
	"todolist/constants"
	"todolist/contract"
	"todolist/domain"
	"todolist/service"
)

func updateHandler(service service.TodoListServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var response contract.TodoListResponse

		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			applogger.Errorf(constants.ErrorFailedToReadRequestBody, "[updateHandler] [ReadAll]", err)
			writeErrorResponse(w, http.StatusBadRequest, err, "updateHandler")
			return
		}

		var task domain.Task

		if err = json.Unmarshal(bytes, &task); err != nil {
			applogger.Errorf(constants.ErrorFailedToUnMarshalRequestBody, "[updateHandler] [Unmarshal]", bytes, err)
			writeErrorResponse(w, http.StatusBadRequest, err, "updateHandler")
			return
		}

		if !task.Validate() {
			applogger.Errorf(constants.ErrorRequestBodyValidationFailed, "[updateHandler] [Validate]")
			writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf(constants.ErrorTaskValidationFailed, "[updateHandler] [Validate]", task), "updateHandler")
			return
		}

		if err = service.Update(&task); err != nil {
			applogger.Errorf(constants.ErrorTaskUpdateFailed, "[updateHandler] [Update]", task, err)
			writeErrorResponse(w, http.StatusInternalServerError, err, "updateHandler")
			return
		}

		response = contract.NewTodoListResponse(nil, constants.EmptyString, true)

		applogger.Infof(constants.SuccessfulAddTask, "[updateHandler]", task)

		writeResponse(w, http.StatusOK, response, "updateHandler")

	}
}

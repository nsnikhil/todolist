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

func removeHandler(service service.TodoListServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var response contract.TodoListResponse

		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			applogger.Errorf(constants.ErrorFailedToReadRequestBody, "[removeHandler] [ReadAll]", err)
			writeErrorResponse(w, http.StatusBadRequest, err, "removeHandler")
			return
		}

		var removeTaskRequest contract.TaskIDRequest

		if err = json.Unmarshal(bytes, &removeTaskRequest); err != nil {
			applogger.Errorf(constants.ErrorFailedToUnMarshalRequestBody, "[removeHandler] [Unmarshal]", bytes, err)
			writeErrorResponse(w, http.StatusBadRequest, err, "removeHandler")
			return
		}

		if !removeTaskRequest.Validate() {
			applogger.Errorf(constants.ErrorRequestBodyValidationFailed, "[removeHandler] [Validate]")
			writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf(constants.ErrorTaskValidationFailed, "[removeHandler] [Validate]", removeTaskRequest.TaskID), "removeHandler")
			return
		}

		if err = service.Remove(removeTaskRequest.TaskID); err != nil {
			applogger.Errorf(constants.ErrorFailedToRemoveTask, "[removeHandler] [Remove]", removeTaskRequest.TaskID, err)
			writeErrorResponse(w, http.StatusInternalServerError, err, "removeHandler")
			return
		}

		response = contract.NewTodoListResponse(nil, constants.EmptyString, true)

		applogger.Infof(constants.SuccessfulRemoveTask, "[removeHandler]", removeTaskRequest.TaskID)

		writeResponse(w, http.StatusOK, response, "removeHandler")

	}
}

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

func addHandler(service service.TodoListServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var response contract.TodoListResponse

		bytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			applogger.Errorf(constants.ErrorFailedToReadRequestBody, "[addHandler] [ReadAll]", err)
			writeErrorResponse(w, http.StatusBadRequest, err, "addHandler")
			return
		}

		var task domain.Task

		if err = json.Unmarshal(bytes, &task); err != nil {
			applogger.Errorf(constants.ErrorFailedToUnMarshalRequestBody, "[addHandler] [Unmarshal]", bytes, err)
			writeErrorResponse(w, http.StatusBadRequest, err, "addHandler")
			return
		}

		if !task.Validate() {
			applogger.Errorf(constants.ErrorRequestBodyValidationFailed, "[addHandler] [Validate]")
			writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf(constants.ErrorTaskValidationFailed, "[addHandler] [Validate]", task), "addHandler")
			return
		}

		if err = service.Add(&task); err != nil {
			applogger.Errorf(constants.ErrorFailedToAddTask, "[addHandler] [Add]", task, err)
			writeErrorResponse(w, http.StatusInternalServerError, err, "addHandler")
			return
		}

		response = contract.NewTodoListResponse(nil, constants.EmptyString, true)

		applogger.Infof(constants.SuccessfulAddTask, "[addHandler]", task)

		writeResponse(w, http.StatusOK, response, "addHandler")

	}
}

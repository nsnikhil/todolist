package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist/applogger"
	"todolist/constants"
	"todolist/contract"
)

func writeResponse(w http.ResponseWriter, statusCode int, body interface{}, handlerName string) {
	w.Header().Set(constants.HeaderContentType, constants.ContentTypeJSON)
	w.WriteHeader(statusCode)

	if body == nil || statusCode == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		applogger.Errorf(constants.ErrorFailedToWriteAPIResponse, fmt.Sprintf("[%s] [writeResponse]", handlerName), err)
	}

}

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error, handlerName string) {
	response := contract.NewTodoListResponse(nil, err.Error(), false)
	writeResponse(w, statusCode, response, handlerName)
}

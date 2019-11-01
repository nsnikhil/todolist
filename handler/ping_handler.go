package handler

import (
	"net/http"
	"todolist/applogger"
	"todolist/constants"
)

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		applogger.Infof(constants.PingHandlerInfo, "[pingHandler]")
		writeResponse(w, http.StatusOK, map[string]interface{}{"message": "pong"}, "pingHandler")
	}
}

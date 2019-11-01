package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ping", nil)

	pingHandler()(w, r)

	expectedResponseBody := "{\"message\":\"pong\"}\n"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponseBody, w.Body.String())
}

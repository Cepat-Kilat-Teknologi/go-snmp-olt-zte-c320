package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func ErrorBadRequest(w http.ResponseWriter, err error) {
	webResponse := ErrorResponse{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: err.Error(),
	}
	SendJSONResponse(w, http.StatusBadRequest, webResponse)
}

func ErrorInternalServerError(w http.ResponseWriter, err error) {
	webResponse := ErrorResponse{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: err.Error(),
	}
	SendJSONResponse(w, http.StatusInternalServerError, webResponse)
}

func ErrorNotFound(w http.ResponseWriter, err error) {
	webResponse := ErrorResponse{
		Code:    http.StatusNotFound,
		Status:  "Not Found",
		Message: err.Error(),
	}
	SendJSONResponse(w, http.StatusNotFound, webResponse)
}

package utils

import (
	"encoding/json"
	"net/http"
)

// ApplicationError will hold the error information
type ApplicationError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// SendSuccess will send the success response to the caller
func SendSuccess(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// SendError will return the error to the user
func SendError(w http.ResponseWriter, status int, errMsg ApplicationError) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errMsg)
}

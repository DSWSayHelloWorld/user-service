package utils

// ApplicationError will hold the error information
type ApplicationError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

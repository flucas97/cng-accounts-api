package errors_utils

import "net/http"

type errorFactory struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestError returns a new bad request as RestErr struct
func NewBadRequestError(message string) *errorFactory {
	return &errorFactory{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError returns a new not found request as RestErr struct
func NewNotFoundError(message string) *errorFactory {
	return &errorFactory{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerError returns a new not found request as RestErr struct
func NewInternalServerError(message string) *errorFactory {
	return &errorFactory{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

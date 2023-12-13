package errors

import "net/http"

type RESTError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RESTError {
	return &RESTError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad request",
	}
}

func NewNotFoundError(message string) *RESTError {
	return &RESTError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewInternalServerError(message string) *RESTError {
	return &RESTError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}

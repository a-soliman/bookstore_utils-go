package rest_errors

import (
	"fmt"
	"net/http"
)

const (
	// DatabaseErrorMsg string
	DatabaseErrorMsg = "database error"
)

// RestErr the restErr interface
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e *restErr) Message() string {
	return e.message
}

func (e *restErr) Status() int {
	return e.status
}

func (e *restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", e.message, e.status, e.error, e.causes)
}

func (e *restErr) Causes() []interface{} {
	return e.causes
}

// NewRestError returns a custom RestError
func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return &restErr{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

// NewBadRequestError returns a RestErr for bad request errors
func NewBadRequestError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

// NewNotFoundError returns a RestErr for not found errors
func NewNotFoundError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

// NewInternalServerError returns a RestErr for internal server error
func NewInternalServerError(message string, err error) RestErr {
	restErr := &restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}
	if err != nil {
		restErr.causes = append(restErr.causes, err.Error())
	}
	return restErr
}

// NewUnAuthorizedError returns a RestError for unauthorized error
func NewUnAuthorizedError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

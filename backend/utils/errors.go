package utils

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Details []string
}

func (e *AppError) Error() string {
	return e.Message
}

// Custom errors
var (
	ErrUnauthorized      = &AppError{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden         = &AppError{Code: http.StatusForbidden, Message: "Forbidden"}
	ErrNotFound          = &AppError{Code: http.StatusNotFound, Message: "Data not found"}
	ErrBadRequest        = &AppError{Code: http.StatusBadRequest, Message: "Bad request"}
	ErrConflict          = &AppError{Code: http.StatusConflict, Message: "Data already exists"}
	ErrUnprocessable     = &AppError{Code: http.StatusUnprocessableEntity, Message: "Invalid input"}
	ErrInternalServer    = &AppError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	ErrInvalidCredentials = &AppError{Code: http.StatusUnauthorized, Message: "Invalid credentials"}
)

func NewAppError(code int, message string, details ...string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

func GetAppError(err error) *AppError {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}
	return ErrInternalServer
}

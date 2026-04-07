package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

// FormatValidationError converts validation errors to user-friendly messages
func FormatValidationError(err error) []string {
	var errMessages []string

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range errs {
			msg := formatFieldError(fieldError)
			errMessages = append(errMessages, msg)
		}
	} else {
		// For non-validator errors (e.g., JSON unmarshal), provide generic message
		errMessages = append(errMessages, "Invalid request format")
	}

	return errMessages
}

// formatFieldError creates a user-friendly error message for a single field
func formatFieldError(fe validator.FieldError) string {
	fieldName := toUserFriendlyFieldName(fe.Field())

	switch fe.Tag() {
	case "required":
		return fieldName + " is required"
	case "email":
		return fieldName + " must be a valid email address"
	case "min":
		return fieldName + " is too short (minimum: " + fe.Param() + ")"
	case "max":
		return fieldName + " is too long (maximum: " + fe.Param() + ")"
	case "oneof":
		return fieldName + " must be one of: " + fe.Param()
	default:
		return fieldName + " is invalid"
	}
}

// toUserFriendlyFieldName converts camelCase to readable format
func toUserFriendlyFieldName(fieldName string) string {
	// Remove common prefixes (PD, RS, etc.)
	if len(fieldName) > 2 && fieldName[0] >= 'A' && fieldName[0] <= 'Z' && fieldName[1] >= 'A' && fieldName[1] <= 'Z' {
		fieldName = fieldName[2:]
	}

	// Insert spaces before capital letters
	var result strings.Builder
	for i, r := range fieldName {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}

	return result.String()
}

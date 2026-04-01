package apperror

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

var (
	ErrNotFound = &AppError{
		Code:    http.StatusNotFound,
		Message: "resource not found",
	}
	ErrInternal = &AppError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}
	ErrExternal = &AppError{
		Code:    http.StatusBadGateway,
		Message: "external service error",
	}
)

func (e *AppError) Wrap(err error) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: e.Message,
		Err:     err,
	}
}

func (e *AppError) WithMessage(msg string) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: msg,
		Err:     e.Err,
	}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func (e *AppError) Unwrap() error {
	return e.Err
}

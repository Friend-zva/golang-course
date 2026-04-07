package domain

import (
	"fmt"
)

type AppStatus int

const (
	RepoNotFound AppStatus = 404
	CodeInternal AppStatus = 500
	CodeExternal AppStatus = 502
)

type AppError struct {
	Code    AppStatus
	Message string
	Err     error
}

var (
	ErrNotFound = &AppError{
		Code:    RepoNotFound,
		Message: "resource not found",
	}
	ErrInternal = &AppError{
		Code:    CodeInternal,
		Message: "internal server error",
	}
	ErrExternal = &AppError{
		Code:    CodeExternal,
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

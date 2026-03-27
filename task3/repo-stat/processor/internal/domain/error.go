package domain

import (
	"fmt"
)

type AppStatus int

const (
	CodeNotFound AppStatus = 404
	CodeInternal AppStatus = 500
	CodeGateway  AppStatus = 502
	CodeTimeout  AppStatus = 504
)

type AppError struct {
	Code    AppStatus
	Message string
	Err     error
}

var (
	ErrNotFound = &AppError{
		Code:    CodeNotFound,
		Message: "resource not found",
	}
	ErrInternal = &AppError{
		Code:    CodeInternal,
		Message: "internal server error",
	}
	ErrGateway = &AppError{
		Code:    CodeGateway,
		Message: "bad gateway: external service error",
	}
	ErrTimeout = &AppError{
		Code:    CodeTimeout,
		Message: "gateway timeout: external service took too long",
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

package domain

import (
	"errors"
	"fmt"
)

type (
	_NotFoundError struct {
		error
	}
	_ValidationError struct {
		error
	}
)

var (
	ErrNotFound   = NotFoundError("")
	ErrValidation = ValidationError("")
)

func NotFoundError(msg string) error {
	if msg == "" {
		msg = "not found"
	}
	return _NotFoundError{errors.New(msg)}
}

func NotFoundErrorf(msgformat string, a ...interface{}) error {
	return NotFoundError(fmt.Sprintf(msgformat, a...))
}

func ValidationError(msg string) error {
	if msg == "" {
		msg = "validation failed"
	}
	return _ValidationError{errors.New(msg)}
}

func ValidationErrorf(msgformat string, a ...interface{}) error {
	return ValidationError(fmt.Sprintf(msgformat, a...))
}

func IsNotFoundError(err error) bool {
	_, ok := err.(_NotFoundError)
	return ok
}

func IsValidationError(err error) bool {
	_, ok := err.(_ValidationError)
	return ok
}

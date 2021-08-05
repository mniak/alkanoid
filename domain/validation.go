package domain

import (
	"strings"

	"github.com/hashicorp/go-multierror"
)

type ValidationResult struct {
	IsValid  bool
	Messages []string
	Error    error
}

func ValidResult() ValidationResult {
	return ValidationResult{
		IsValid:  true,
		Messages: []string{},
	}
}

func InvalidResult(message string) ValidationResult {
	return ValidResult().AppendMessage(message)
}

func ValidationErrorResult(err error) ValidationResult {
	return ValidResult().AppendError(err)
}

func (vr ValidationResult) AppendMessage(message string) ValidationResult {
	if message != "" {
		vr.Messages = append(vr.Messages, message)
		vr.IsValid = false
	}
	return vr
}

func (vr ValidationResult) AppendError(err error) ValidationResult {
	if err != nil {
		vr.Error = multierror.Append(vr.Error, err)
		vr.IsValid = false
	}
	return vr
}

func (vr ValidationResult) String() string {
	return strings.Join(vr.Messages, ". ")
}

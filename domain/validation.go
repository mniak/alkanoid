package domain

import (
	"fmt"
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

func (vr ValidationResult) AppendMessage(messageOrFormat string, args ...interface{}) ValidationResult {
	if messageOrFormat != "" {
		if len(args) > 0 {
			messageOrFormat = fmt.Sprintf(messageOrFormat, args...)
		}
		vr.Messages = append(vr.Messages, messageOrFormat)
		vr.IsValid = false
	}
	return vr
}

func (vr ValidationResult) Combine(vr2 ValidationResult) ValidationResult {
	for _, msg := range vr2.Messages {
		vr = vr.AppendMessage(msg)
	}

	vr = vr.AppendError(vr2.Error)
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

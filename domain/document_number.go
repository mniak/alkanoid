package domain

import (
	"regexp"
)

type DocumentNumber string

func (dn DocumentNumber) Validate() ValidationResult {
	if dn == "" {
		return InvalidResult("empty document number")
	}

	ok, err := regexp.MatchString(`^[0-9/.\-]+$`, string(dn))
	if err != nil {
		return ValidationErrorResult(err)
	}
	if !ok {
		return InvalidResult("invalid characters on document number")
	}

	return ValidResult()
}

func (dn DocumentNumber) String() string {
	return string(dn)
}

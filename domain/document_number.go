package domain

import (
	"regexp"
)

type DocumentNumber string

func (dn DocumentNumber) Validate() ValidationResult {
	if dn == "" {
		return InvalidResult("empty document number")
	}

	if len(dn) > 18 {
		return InvalidResult("document number is too lengthy")
	}

	ok, err := regexp.MatchString(`^[0-9/.\-]+$`, string(dn))
	if err != nil {
		return ValidationErrorResult(err)
	}
	if !ok {
		return InvalidResult("document number contains invalid characters")
	}

	return ValidResult()
}

func (dn DocumentNumber) String() string {
	return string(dn)
}

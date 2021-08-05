package domain

import (
	"errors"
	"regexp"
)

type DocumentNumber string

var ErrInvalidDocumentNumber = errors.New("invalid document number")

func (dn DocumentNumber) Validate() error {
	if dn == "" {
		return ErrInvalidDocumentNumber
	}

	ok, err := regexp.MatchString(`^[0-9/.\-]+$`, string(dn))
	if err != nil {
		return err
	}
	if !ok {
		return ErrInvalidDocumentNumber
	}

	return nil
}

func (dn DocumentNumber) String() string {
	return string(dn)
}

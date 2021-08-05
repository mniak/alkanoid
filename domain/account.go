package domain

import (
	"github.com/hashicorp/go-multierror"
)

type Account struct {
	ID             int
	DocumentNumber DocumentNumber
}

func NewAccount(doc DocumentNumber) Account {
	return Account{
		DocumentNumber: doc,
	}
}

func (a Account) Validate() error {
	var err error
	err = multierror.Append(err, a.DocumentNumber.Validate())
	return err
}

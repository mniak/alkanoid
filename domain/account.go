package domain

import (
	"github.com/hashicorp/go-multierror"
)

type Account struct {
	id  int
	doc DocumentNumber
}

func NewAccount(doc DocumentNumber) Account {
	return Account{
		doc: doc,
	}
}

func ExistingAccount(id int, doc DocumentNumber) Account {
	return Account{
		id:  id,
		doc: doc,
	}
}

func (a Account) Validate() error {
	var err error
	err = multierror.Append(err, a.doc.Validate())
	return err
}

func (a Account) ID() int {
	return a.id
}

func (a Account) DocumentNumber() DocumentNumber {
	return a.doc
}

func (a Account) WithID(id int) Account {
	a.id = id
	return a
}

package domain

import "errors"

type (
	AccountRepository interface {
		Save(acc Account) (int, error)
		Load(id int) (Account, error)
	}
	TransactionRepository interface {
		Save(tra Transaction) (int, error)
		Load(id int) (Transaction, error)
	}
)

var (
	ErrAccountNotFound     error = errors.New("account not found")
	ErrTransactionNotFound error = errors.New("transaction not found")
)

package domain

import (
	"time"

	"github.com/hashicorp/go-multierror"
)

type Transaction struct {
	ID            int
	AccountID     int
	OperationType OperationType
	Amount        float64
	EventDate     time.Time
}

func NewTransaction(
	accountID int,
	operationType OperationType,
	amount float64,
) Transaction {
	return Transaction{
		AccountID:     accountID,
		OperationType: operationType,
		Amount:        amount,
		EventDate:     time.Now().UTC(),
	}
}

func (t Transaction) Validate() error {
	var err error
	err = multierror.Append(err, t.OperationType.Validate())
	return err
}

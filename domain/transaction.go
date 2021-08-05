package domain

import (
	"time"
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

func (t Transaction) Validate() ValidationResult {
	return t.OperationType.Validate()
}

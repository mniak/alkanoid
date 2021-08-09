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
	result := t.OperationType.Validate()

	if t.OperationType == OpCompraAVista ||
		t.OperationType == OpCompraParcelada ||
		t.OperationType == OpSaque {
		if t.Amount > 0 {
			result = result.AppendMessage("operations of type %s require a negative amount", t.OperationType)
		}
	} else if t.OperationType == OpPagamento {
		if t.Amount < 0 {
			result = result.AppendMessage("operations of type %s require a positive amount", t.OperationType)
		}
	}
	return result
}

package app

import "github.com/mniak/Alkanoid/domain"

type Mapper interface {
	AccountFromCreateAccountRequest(req CreateAccountRequest) domain.Account
	GetAccountResponseFromAccount(acc domain.Account) GetAccountResponse
	TransactionFromCreateTransactionRequest(req CreateTransactionRequest) domain.Transaction
}
type _Mapper struct{}

func NewMapper() Mapper {
	return &_Mapper{}
}

func (_ _Mapper) AccountFromCreateAccountRequest(req CreateAccountRequest) domain.Account {
	account := domain.NewAccount(
		domain.DocumentNumber(req.DocumentNumber),
	)
	return account
}

func (_ _Mapper) GetAccountResponseFromAccount(acc domain.Account) GetAccountResponse {
	return GetAccountResponse{
		AccountID:            acc.ID,
		DocumentNumber:       acc.DocumentNumber.String(),
		AvailableCreditLimit: acc.AvailableCreditLimit,
	}
}

func (_ _Mapper) TransactionFromCreateTransactionRequest(req CreateTransactionRequest) domain.Transaction {
	optype := domain.OperationType(req.OperationTypeID)

	var sign int
	if optype.IsDeposit() {
		sign = 1
	} else {
		sign = -1
	}

	transaction := domain.NewTransaction(
		req.AccountID,
		optype,
		req.Amount*float64(sign),
	)
	return transaction
}

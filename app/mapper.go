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
		AccountID:      acc.ID,
		DocumentNumber: acc.DocumentNumber.String(),
	}
}

func (_ _Mapper) TransactionFromCreateTransactionRequest(req CreateTransactionRequest) domain.Transaction {
	transaction := domain.NewTransaction(
		req.AccountID,
		domain.OperationType(req.OperationTypeID),
		req.Amount,
	)
	return transaction
}

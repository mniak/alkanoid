package app

import (
	"github.com/mniak/Alkanoid/domain"
)

type _Application struct {
	repos RepositoriesRegistry
}

func NewApplicationWithoutDecorators(repos RepositoriesRegistry) _Application {
	return _Application{repos: repos}
}

func NewApplication(repos RepositoriesRegistry) _Application {
	a := NewApplicationWithoutDecorators(repos)
	a.repos = a.repos.withValidation(newValidationServicesFromApp(a))
	return a
}

func (a _Application) CreateAccount(req CreateAccountRequest) (resp CreateAccountResponse, err error) {
	account := domain.NewAccount(
		domain.DocumentNumber(req.DocumentNumber),
	)
	id, err := a.repos.Account.Save(account)
	if err != nil {
		return
	}

	resp.AccountID = id
	return
}

func (a _Application) GetAccount(req GetAccountRequest) (resp GetAccountResponse, err error) {
	acc, err := a.repos.Account.Load(req.AccountID)
	if err != nil {
		return
	}

	resp.AccountID = acc.ID
	resp.DocumentNumber = acc.DocumentNumber.String()

	return
}

func (a _Application) CreateTransaction(req CreateTransactionRequest) (resp CreateTransactionResponse, err error) {
	transaction := domain.NewTransaction(
		req.AccountID,
		domain.OperationType(req.OperationTypeID),
		req.Amount,
	)
	id, err := a.repos.Transaction.Save(transaction)
	if err != nil {
		return
	}

	resp.TransactionID = id
	return
}

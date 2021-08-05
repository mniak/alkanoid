package application

import "github.com/mniak/Alkanoid/domain"

type _Application struct {
	accountRepo     domain.AccountRepository
	transactionRepo domain.TransactionRepository
}

func New(
	accountRepo domain.AccountRepository,
	transactionRepo domain.TransactionRepository,
) _Application {
	return _Application{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (a _Application) CreateAccount(req CreateAccountRequest) (resp CreateAccountResponse, err error) {
	account := domain.NewAccount(
		domain.DocumentNumber(req.DocumentNumber),
	)
	err = account.Validate()
	if err != nil {
		return
	}

	id, err := a.accountRepo.Save(account)
	if err != nil {
		return
	}
	resp.AccountID = id
	return
}

func (a _Application) GetAccount(req GetAccountRequest) (GetAccountResponse, error) {
	return GetAccountResponse{}, nil
}

func (a _Application) CreateTransaction(req CreateTransactionRequest) (resp CreateTransactionResponse, err error) {
	transaction := domain.NewTransaction(
		req.AccountID,
		domain.OperationType(req.OperationTypeID),
		req.Amount,
	)
	err = transaction.Validate()
	if err != nil {
		return
	}

	id, err := a.transactionRepo.Save(transaction)
	if err != nil {
		return
	}

	resp.TransactionID = id
	return
}

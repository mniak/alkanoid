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

func (a _Application) CreateAccount(req CreateAccountRequest) (CreateAccountResponse, error) {
	account := domain.NewAccount(
		domain.DocumentNumber(req.DocumentNumber),
	)
	err := account.Validate()
	if err != nil {
		return CreateAccountResponse{}, err
	}

	id, err := a.accountRepo.Save(account)
	if err != nil {
		return CreateAccountResponse{}, err
	}

	return CreateAccountResponse{
		AccountID: id,
	}, nil
}

func (a _Application) GetAccount(GetAccountRequest) (GetAccountResponse, error) {
	return GetAccountResponse{}, nil
}

func (a _Application) CreateTransaction(CreateTransactionRequest) (CreateTransactionResponse, error) {
	return CreateTransactionResponse{}, nil
}

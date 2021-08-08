package app

import "github.com/mniak/Alkanoid/domain"

type ValidationServicesRegistry struct {
	Account     domain.AccountValidationService
	Transaction domain.TransactionValidationService
}

func newValidationServicesFromApp(a _Application) ValidationServicesRegistry {
	return ValidationServicesRegistry{
		Account:     domain.NewAccountValidationService(),
		Transaction: domain.NewTransactionValidationService(a.repos.Account),
	}
}

type RepositoriesRegistry struct {
	Account     domain.AccountRepository
	Transaction domain.TransactionRepository
}

func (rr RepositoriesRegistry) withValidation(vsr ValidationServicesRegistry) RepositoriesRegistry {
	rr.Account = WrapAccountRepoWithValidation(rr.Account, vsr.Account)
	rr.Transaction = WrapTransactionRepoWithValidation(rr.Transaction, vsr.Transaction)
	return rr
}

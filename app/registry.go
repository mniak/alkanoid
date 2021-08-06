package app

import "github.com/mniak/Alkanoid/domain"

type ValidationServicesRegistry struct {
	Transaction domain.TransactionValidationService
}

func newValidationServicesFromApp(a _Application) ValidationServicesRegistry {
	return ValidationServicesRegistry{
		Transaction: domain.NewTransactionValidationService(a.repos.Account),
	}
}

type RepositoriesRegistry struct {
	Account     domain.AccountRepository
	Transaction domain.TransactionRepository
}

func (rr RepositoriesRegistry) withValidation(vsr ValidationServicesRegistry) RepositoriesRegistry {
	rr.Transaction = WrapTransactionRepoWithValidation(rr.Transaction, vsr.Transaction)
	return rr
}

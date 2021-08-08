package app

import (
	"github.com/mniak/Alkanoid/domain"
)

type _AccountRepoWithValidation struct {
	domain.AccountRepository
	accountValidationService domain.AccountValidationService
}

func WrapAccountRepoWithValidation(inner domain.AccountRepository, accountValidationService domain.AccountValidationService) domain.AccountRepository {
	return &_AccountRepoWithValidation{
		AccountRepository:        inner,
		accountValidationService: accountValidationService,
	}
}

func (repo *_AccountRepoWithValidation) Save(trans domain.Account) (int, error) {
	valres := repo.accountValidationService.Validate(trans)
	if valres.Error != nil {
		return 0, valres.Error
	} else if !valres.IsValid {
		return 0, domain.ValidationErrorf("cannot create transaction: %s", valres)
	}

	return repo.AccountRepository.Save(trans)
}

type _TransactionRepoWithValidation struct {
	domain.TransactionRepository
	transactionValidationService domain.TransactionValidationService
}

func WrapTransactionRepoWithValidation(inner domain.TransactionRepository, transactionValidationService domain.TransactionValidationService) domain.TransactionRepository {
	return &_TransactionRepoWithValidation{
		TransactionRepository:        inner,
		transactionValidationService: transactionValidationService,
	}
}

func (repo *_TransactionRepoWithValidation) Save(trans domain.Transaction) (int, error) {
	valres := repo.transactionValidationService.Validate(trans)
	if valres.Error != nil {
		return 0, valres.Error
	} else if !valres.IsValid {
		return 0, domain.ValidationErrorf("cannot create transaction: %s", valres)
	}

	return repo.TransactionRepository.Save(trans)
}

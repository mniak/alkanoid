package app

import (
	"github.com/mniak/Alkanoid/domain"
)

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

func (tr *_TransactionRepoWithValidation) Save(trans domain.Transaction) (int, error) {
	valres := tr.transactionValidationService.Validate(trans)
	if valres.Error != nil {
		return 0, valres.Error
	} else if !valres.IsValid {
		return 0, domain.ValidationErrorf("cannot create transaction: %s", valres)
	}

	return tr.TransactionRepository.Save(trans)
}

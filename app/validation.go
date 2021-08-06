package app

import (
	"github.com/mniak/Alkanoid/domain"
)

type _TransactionRepoWithValidation struct {
	domain.TransactionRepository
	validationServices ValidationServicesRegistry
}

func WrapTransactionRepoWithValidation(inner domain.TransactionRepository, validationServices ValidationServicesRegistry) domain.TransactionRepository {
	return &_TransactionRepoWithValidation{
		TransactionRepository: inner,
		validationServices:    validationServices,
	}
}

func (tr *_TransactionRepoWithValidation) Save(trans domain.Transaction) (int, error) {
	valres := tr.validationServices.Transaction.Validate(trans)
	if valres.Error != nil {
		return 0, valres.Error
	} else if !valres.IsValid {
		return 0, domain.ValidationErrorf("cannot create transaction: %s", valres)
	}

	return tr.TransactionRepository.Save(trans)
}

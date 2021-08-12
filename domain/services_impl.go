package domain

import "fmt"

type _AccountValidationService struct{}

func NewAccountValidationService() AccountValidationService {
	return &_AccountValidationService{}
}

func (s *_AccountValidationService) Validate(a Account) ValidationResult {
	result := ValidResult()
	result = result.Combine(a.Validate())
	return result
}

type _TransactionValidationService struct {
	acctRepo AccountRepository
}

func NewTransactionValidationService(acctRepo AccountRepository) TransactionValidationService {
	return &_TransactionValidationService{
		acctRepo: acctRepo,
	}
}

func (s *_TransactionValidationService) Validate(t Transaction) ValidationResult {
	result := t.Validate()
	exists, err := s.acctRepo.Exists(t.AccountID)
	result = result.AppendError(err)
	if err == nil && !exists {
		result = result.AppendMessage(fmt.Sprintf("account with id %d not found", t.AccountID))
	}

	/*

			se for de gastar então:
		       limite = pegalimite(conta_id)
			   se limite < trans.valor {
				   falhar
			   } senão {
				   blz
				   e descontar do limite
			   }
			senão (se de depositar) então:
			   aumenta limite
			fim
	*/
	return result
}

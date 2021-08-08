package domain

type AccountValidationService interface {
	Validate(t Account) ValidationResult
}

type TransactionValidationService interface {
	Validate(t Transaction) ValidationResult
}

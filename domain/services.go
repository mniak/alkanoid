package domain

type TransactionValidationService interface {
	Validate(t Transaction) ValidationResult
}

package alkanoid

//go:generate mockgen -package=mocks -destination=internal/mocks/app.go	github.com/mniak/Alkanoid/app Application

//go:generate mockgen -package=mocks -destination=internal/mocks/transaction_repository.go	github.com/mniak/Alkanoid/domain TransactionRepository
//go:generate mockgen -package=mocks -destination=internal/mocks/transaction_validation_service.go	github.com/mniak/Alkanoid/domain TransactionValidationService

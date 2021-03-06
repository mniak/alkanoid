package alkanoid

//go:generate mockgen -package=mocks -destination=internal/mocks/app.go		github.com/mniak/Alkanoid/app Application
//go:generate mockgen -package=mocks -destination=internal/mocks/mapper.go	github.com/mniak/Alkanoid/app Mapper

//go:generate mockgen -package=mocks -destination=internal/mocks/account_repository.go			github.com/mniak/Alkanoid/domain AccountRepository
//go:generate mockgen -package=mocks -destination=internal/mocks/account_validation_service.go	github.com/mniak/Alkanoid/domain AccountValidationService

//go:generate mockgen -package=mocks -destination=internal/mocks/transaction_repository.go			github.com/mniak/Alkanoid/domain TransactionRepository
//go:generate mockgen -package=mocks -destination=internal/mocks/transaction_validation_service.go	github.com/mniak/Alkanoid/domain TransactionValidationService

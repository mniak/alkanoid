package application

type Application interface {
	CreateAccount(CreateAccountRequest) (CreateAccountResponse, error)
	GetAccount(GetAccountRequest) (GetAccountResponse, error)
	CreateTransaction(CreateTransactionRequest) (CreateTransactionResponse, error)
}

package application

type _Application struct{}

func New() _Application {
	return _Application{}
}

func (a _Application) CreateAccount(CreateAccountRequest) (CreateAccountResponse, error) {
	return CreateAccountResponse{}, nil
}

func (a _Application) GetAccount(GetAccountRequest) (GetAccountResponse, error) {
	return GetAccountResponse{}, nil
}

func (a _Application) CreateTransaction(CreateTransactionRequest) (CreateTransactionResponse, error) {
	return CreateTransactionResponse{}, nil
}

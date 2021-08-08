package app

type _Application struct {
	repos  RepositoriesRegistry
	mapper Mapper
}

func NewApplicationWithoutMagic(repos RepositoriesRegistry, mapper Mapper) _Application {
	return _Application{
		repos:  repos,
		mapper: mapper,
	}
}

func NewApplication(repos RepositoriesRegistry) _Application {
	mapper := NewMapper()
	a := NewApplicationWithoutMagic(repos, mapper)
	a.repos = a.repos.withValidation(newValidationServicesFromApp(a))
	return a
}

func (a _Application) CreateAccount(req CreateAccountRequest) (resp CreateAccountResponse, err error) {
	account := a.mapper.AccountFromCreateAccountRequest(req)
	id, err := a.repos.Account.Save(account)
	if err != nil {
		return
	}

	resp.AccountID = id
	return
}

func (a _Application) GetAccount(req GetAccountRequest) (GetAccountResponse, error) {
	acc, err := a.repos.Account.Load(req.AccountID)
	if err != nil {
		return GetAccountResponse{}, err
	}

	return a.mapper.GetAccountResponseFromAccount(acc), nil
}

func (a _Application) CreateTransaction(req CreateTransactionRequest) (resp CreateTransactionResponse, err error) {
	transaction := a.mapper.TransactionFromCreateTransactionRequest(req)
	id, err := a.repos.Transaction.Save(transaction)
	if err != nil {
		return
	}

	resp.TransactionID = id
	return
}

package domain

type Account struct {
	ID             int
	DocumentNumber DocumentNumber
}

func NewAccount(doc DocumentNumber) Account {
	return Account{
		DocumentNumber: doc,
	}
}

func (a Account) Validate() ValidationResult {
	return a.DocumentNumber.Validate()
}

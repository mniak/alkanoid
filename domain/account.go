package domain

type Account struct {
	ID                   int
	DocumentNumber       DocumentNumber
	AvailableCreditLimit float64
}

func NewAccount(doc DocumentNumber) Account {
	return Account{
		DocumentNumber: doc,
	}
}

func (a Account) Validate() ValidationResult {
	result := ValidResult()
	if a.AvailableCreditLimit < 0 {
		result = result.AppendMessage("credit limit cannot be negative")
	}
	result = result.Combine(a.DocumentNumber.Validate())
	return result
}

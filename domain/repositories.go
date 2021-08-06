package domain

type (
	AccountRepository interface {
		Save(acc Account) (int, error)
		Load(id int) (Account, error)
		Exists(id int) (bool, error)
	}
	TransactionRepository interface {
		Save(tra Transaction) (int, error)
		Load(id int) (Transaction, error)
	}
)

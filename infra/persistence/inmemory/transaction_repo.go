package inmemory

import (
	"github.com/mniak/Alkanoid/domain"
)

type _TransactionRepository struct {
	data  map[int]domain.Transaction
	maxId int
}

func NewTransactionRepository() domain.TransactionRepository {
	return &_TransactionRepository{
		data: make(map[int]domain.Transaction),
	}
}

func (r *_TransactionRepository) Save(acc domain.Transaction) (int, error) {
	if acc.ID == 0 {
		r.maxId++
		acc.ID = r.maxId
	} else if acc.ID > r.maxId {
		r.maxId = acc.ID
	}
	r.data[acc.ID] = acc
	return acc.ID, nil
}

func (r *_TransactionRepository) Load(id int) (domain.Transaction, error) {
	acc, ok := r.data[id]
	if !ok {
		return domain.Transaction{}, domain.ErrNotFound
	}
	return acc, nil
}

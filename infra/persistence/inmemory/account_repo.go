package inmemory

import (
	"github.com/mniak/Alkanoid/domain"
)

type _AccountRepository struct {
	data  map[int]domain.Account
	maxId int
}

func NewAccountRepository() domain.AccountRepository {
	return &_AccountRepository{
		data: make(map[int]domain.Account),
	}
}

func (r *_AccountRepository) Save(acc domain.Account) (int, error) {
	if acc.ID == 0 {
		r.maxId++
		acc.ID = r.maxId
	} else if acc.ID > r.maxId {
		r.maxId = acc.ID
	}
	r.data[acc.ID] = acc
	return acc.ID, nil
}

func (r *_AccountRepository) Load(id int) (domain.Account, error) {
	acc, ok := r.data[id]
	if !ok {
		return domain.Account{}, domain.ErrAccountNotFound
	}
	return acc, nil
}

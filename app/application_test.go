package app_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/domain"
	"github.com/mniak/Alkanoid/internal/matchers"
	"github.com/mniak/Alkanoid/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplication_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var req app.CreateAccountRequest
	gofakeit.Struct(&req)
	id := int(gofakeit.Int32())

	acctRepo := mocks.NewMockAccountRepository(ctrl)
	acctRepo.EXPECT().Save(domain.Account{
		DocumentNumber: domain.DocumentNumber(req.DocumentNumber),
	}).Return(id, nil)

	sut := app.NewApplicationWithoutMagic(app.RepositoriesRegistry{
		Account: acctRepo,
	})

	resp, err := sut.CreateAccount(req)
	require.NoError(t, err)
	assert.Equal(t, id, resp.AccountID)
}

func TestApplication_GetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var req app.GetAccountRequest
	gofakeit.Struct(&req)
	var acct domain.Account
	gofakeit.Struct(&acct)

	acctRepo := mocks.NewMockAccountRepository(ctrl)
	acctRepo.EXPECT().Load(req.AccountID).Return(acct, nil)

	sut := app.NewApplicationWithoutMagic(app.RepositoriesRegistry{
		Account: acctRepo,
	})

	resp, err := sut.GetAccount(req)
	require.NoError(t, err)
	assert.Equal(t, id, resp.AccountID)
}

func TestApplication_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var req app.CreateTransactionRequest
	gofakeit.Struct(&req)
	id := int(gofakeit.Int32())

	tranRepo := mocks.NewMockTransactionRepository(ctrl)
	tranRepo.EXPECT().Save(gomock.All(
		matchers.TransactionFieldEquals("AccountID", req.AccountID, func(tr domain.Transaction) interface{} {
			return tr.AccountID
		}),
		matchers.TransactionFieldEquals("Amount", req.Amount, func(tr domain.Transaction) interface{} {
			return tr.Amount
		}),
		matchers.TransactionFieldEquals("OperationType", req.OperationTypeID, func(tr domain.Transaction) interface{} {
			return tr.OperationType.ID()
		}),
	)).Return(id, nil)

	sut := app.NewApplicationWithoutMagic(app.RepositoriesRegistry{
		Transaction: tranRepo,
	})

	resp, err := sut.CreateTransaction(req)
	require.NoError(t, err)
	assert.Equal(t, id, resp.TransactionID)
}

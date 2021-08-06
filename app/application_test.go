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

	sut := app.NewApplicationWithoutDecorators(app.RepositoriesRegistry{
		Account: acctRepo,
	})

	resp, err := sut.CreateAccount(req)
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
	tranRepo.EXPECT().Save(matchers.InlineMatcher(func(x interface{}) bool {
		return false
	})).Return(id, nil)

	sut := app.NewApplicationWithoutDecorators(app.RepositoriesRegistry{
		Transaction: tranRepo,
	})

	resp, err := sut.CreateTransaction(req)
	require.NoError(t, err)
	assert.Equal(t, id, resp.TransactionID)
}

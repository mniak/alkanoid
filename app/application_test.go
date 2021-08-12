package app_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/domain"
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

	acc := domain.Account{
		DocumentNumber: domain.DocumentNumber(req.DocumentNumber),
	}
	mapper := mocks.NewMockMapper(ctrl)
	mapper.EXPECT().AccountFromCreateAccountRequest(req).Return(acc)
	acctRepo := mocks.NewMockAccountRepository(ctrl)
	acctRepo.EXPECT().Save(acc).Return(id, nil)

	sut := app.NewApplicationWithoutMagic(app.RepositoriesRegistry{
		Account: acctRepo,
	}, mapper)

	resp, err := sut.CreateAccount(req)
	require.NoError(t, err)
	assert.Equal(t, id, resp.AccountID)
}

func TestApplication_GetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var req app.GetAccountRequest
	gofakeit.Struct(&req)
	var acc domain.Account
	gofakeit.Struct(&acc)
	var expectedResp app.GetAccountResponse
	gofakeit.Struct(&expectedResp)

	acctRepo := mocks.NewMockAccountRepository(ctrl)
	acctRepo.EXPECT().Load(req.AccountID).Return(acc, nil)
	mapper := mocks.NewMockMapper(ctrl)
	mapper.EXPECT().GetAccountResponseFromAccount(acc).Return(expectedResp)

	sut := app.NewApplicationWithoutMagic(app.RepositoriesRegistry{
		Account: acctRepo,
	}, mapper)

	resp, err := sut.GetAccount(req)
	require.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestApplication_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var req app.CreateTransactionRequest
	gofakeit.Struct(&req)
	id := int(gofakeit.Int32())

	var tran domain.Transaction
	gofakeit.Struct(&tran)
	tran.Amount = float64(gofakeit.Number(10, 100))

	var acct domain.Account
	gofakeit.Struct(&acct)
	acct.AvailableCreditLimit = float64(gofakeit.Number(150, 200))

	mapper := mocks.NewMockMapper(ctrl)
	mapper.EXPECT().TransactionFromCreateTransactionRequest(req).Return(tran)

	tranRepo := mocks.NewMockTransactionRepository(ctrl)
	tranRepo.EXPECT().Save(tran).Return(id, nil)

	acctRepo := mocks.NewMockAccountRepository(ctrl)
	acctRepo.EXPECT().Load(req.AccountID).Return(acct, nil)
	acctRepo.EXPECT().Save(gomock.Any()).Return(acct.ID, nil)

	sut := app.NewApplicationWithoutMagic(app.RepositoriesRegistry{
		Account:     acctRepo,
		Transaction: tranRepo,
	}, mapper)

	resp, err := sut.CreateTransaction(req)
	require.NoError(t, err)
	assert.Equal(t, id, resp.TransactionID)
}

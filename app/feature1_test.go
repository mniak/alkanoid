package app_test

import (
	"testing"

	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/domain"
	"github.com/mniak/Alkanoid/infra/persistence/inmemory"
	"github.com/stretchr/testify/require"
)

func TestFeature1_Example1(t *testing.T) {
	a := app.NewApplication(
		app.RepositoriesRegistry{
			Account:     inmemory.NewAccountRepository(),
			Transaction: inmemory.NewTransactionRepository(),
		},
	)

	accResp, err := a.CreateAccount(app.CreateAccountRequest{
		DocumentNumber: "12345678900",
	})
	require.NoError(t, err)

	_, err = a.CreateTransaction(app.CreateTransactionRequest{
		AccountID:       accResp.AccountID,
		OperationTypeID: domain.OpPagamento,
		Amount:          100,
	})
	require.NoError(t, err)

	acc100, err := a.GetAccount(app.GetAccountRequest{
		AccountID: accResp.AccountID,
	})
	require.NoError(t, err)
	require.Equal(t, float64(100), acc100.AvailableCreditLimit)

	_, err = a.CreateTransaction(app.CreateTransactionRequest{
		AccountID:       accResp.AccountID,
		OperationTypeID: domain.OpSaque,
		Amount:          30,
	})
	require.NoError(t, err)

	acc70, err := a.GetAccount(app.GetAccountRequest{
		AccountID: accResp.AccountID,
	})
	require.NoError(t, err)
	require.Equal(t, float64(70), acc70.AvailableCreditLimit)

	_, err = a.CreateTransaction(app.CreateTransactionRequest{
		AccountID:       accResp.AccountID,
		OperationTypeID: domain.OpSaque,
		Amount:          80,
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "credit limit cannot be negative")

	_, err = a.CreateTransaction(app.CreateTransactionRequest{
		AccountID:       accResp.AccountID,
		OperationTypeID: domain.OpPagamento,
		Amount:          20,
	})
	require.NoError(t, err)

	acc90, err := a.GetAccount(app.GetAccountRequest{
		AccountID: accResp.AccountID,
	})
	require.NoError(t, err)
	require.Equal(t, float64(90), acc90.AvailableCreditLimit)
}

package app_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/domain"
	"github.com/stretchr/testify/assert"
)

func TestMapperAccountFromCreateAccountRequest(t *testing.T) {
	var req app.CreateAccountRequest
	gofakeit.Struct(&req)

	sut := app.NewMapper()
	acc := sut.AccountFromCreateAccountRequest(req)

	assert.Equal(t, req.DocumentNumber, acc.DocumentNumber.String())
}

func TestGetAccountResponseFromAccount(t *testing.T) {
	var acc domain.Account
	gofakeit.Struct(&acc)

	sut := app.NewMapper()
	resp := sut.GetAccountResponseFromAccount(acc)

	assert.Equal(t, acc.ID, resp.AccountID)
	assert.Equal(t, acc.DocumentNumber.String(), resp.DocumentNumber)
}

func TestTransactionFromCreateTransactionRequest(t *testing.T) {
	var req app.CreateTransactionRequest
	gofakeit.Struct(&req)

	sut := app.NewMapper()
	tran := sut.TransactionFromCreateTransactionRequest(req)

	assert.Equal(t, tran.AccountID, req.AccountID)
	assert.Equal(t, tran.OperationType.ID(), req.OperationTypeID)
	assert.Equal(t, tran.Amount, req.Amount)
	assert.InDelta(t, time.Now().UnixNano(), tran.EventDate.UnixNano(), float64((100 * time.Millisecond).Nanoseconds()))
}

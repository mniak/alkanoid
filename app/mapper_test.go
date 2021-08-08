package app_test

import (
	"testing"

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

	assert.Equal(t, resp.AccountID, acc.ID)
	assert.Equal(t, resp.DocumentNumber, acc.DocumentNumber.String())
}

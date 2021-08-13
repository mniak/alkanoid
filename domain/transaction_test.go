package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionValidate_WhenCompraOrSaque_AmountShouldBeNegativeOr0(t *testing.T) {
	testCases := []struct {
		optype  OperationType
		message string
	}{
		{
			optype:  OpCompraAVista,
			message: "invalid amout for operation of type COMPRA A VISTA",
		},
		{
			optype:  OpCompraParcelada,
			message: "invalid amout for operation of type COMPRA PARCELADA",
		},
		{
			optype:  OpSaque,
			message: "invalid amout for operation of type SAQUE",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.optype.String(), func(t *testing.T) {
			tran := NewTransaction(123, tc.optype, -5)
			valres := tran.Validate()
			assert.True(t, valres.IsValid)

			tran = NewTransaction(123, tc.optype, 0)
			valres = tran.Validate()
			assert.True(t, valres.IsValid)

			tran = NewTransaction(123, tc.optype, 127)
			valres = tran.Validate()
			assert.False(t, valres.IsValid)
			assert.Contains(t, valres.Messages, tc.message)
		})
	}
}

func TestTransactionValidate_WhenPagamento_AmountShouldBePositiveOr0(t *testing.T) {
	testCases := []struct {
		optype  OperationType
		message string
	}{
		{
			optype:  OpPagamento,
			message: "invalid amout for operation of type PAGAMENTO",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.optype.String(), func(t *testing.T) {
			tran := NewTransaction(123, tc.optype, -5)
			valres := tran.Validate()
			assert.False(t, valres.IsValid)
			assert.Contains(t, valres.Messages, tc.message)

			tran = NewTransaction(123, tc.optype, 0)
			valres = tran.Validate()
			assert.True(t, valres.IsValid)

			tran = NewTransaction(123, tc.optype, 127)
			valres = tran.Validate()
			assert.True(t, valres.IsValid)
		})
	}
}

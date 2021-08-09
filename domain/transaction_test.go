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
			message: "operations of type COMPRA A VISTA require a negative amount",
		},
		{
			optype:  OpCompraParcelada,
			message: "operations of type COMPRA PARCELADA require a negative amount",
		},
		{
			optype:  OpSaque,
			message: "operations of type SAQUE require a negative amount",
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
			message: "operations of type PAGAMENTO require a positive amount",
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

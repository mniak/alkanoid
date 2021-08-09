package domain

import "fmt"

type OperationType int

const (
	OpCompraAVista    = 1
	OpCompraParcelada = 2
	OpSaque           = 3
	OpPagamento       = 4
)

func (ot OperationType) String() string {
	switch ot {
	case OpCompraAVista:
		return "COMPRA A VISTA"
	case OpCompraParcelada:
		return "COMPRA PARCELADA"
	case OpSaque:
		return "SAQUE"
	case OpPagamento:
		return "PAGAMENTO"
	}
	return "TIPO DE OPERACAO INVALIDO"
}

func (ot OperationType) Validate() ValidationResult {
	if ot < OpCompraAVista || ot > OpPagamento {
		return InvalidResult(fmt.Sprintf("invalid operation type %d", int(ot)))
	}
	return ValidResult()
}

func (ot OperationType) ID() int {
	return int(ot)
}

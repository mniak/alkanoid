package domain

import "fmt"

type OperationType int

const (
	CompraAVista    = 1
	CompraParcelada = 2
	Saque           = 3
	Pagamento       = 4
)

func (ot OperationType) String() string {
	switch ot {
	case CompraAVista:
		return "COMPRA A VISTA"
	case CompraParcelada:
		return "COMPRA PARCELADA"
	case Saque:
		return "SAQUE"
	case Pagamento:
		return "PAGAMENTO"
	}
	return "TIPO DE OPERACAO INVALIDO"
}

func (ot OperationType) Validate() ValidationResult {
	if ot < CompraAVista || ot > Pagamento {
		return InvalidResult(fmt.Sprintf("invalid operation type %d", int(ot)))
	}
	return ValidResult()
}

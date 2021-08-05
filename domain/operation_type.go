package domain

import "errors"

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

var ErrInvalidOperationType error = errors.New("invalid operation type")

func (ot OperationType) Validate() error {
	if ot < CompraAVista || ot > Pagamento {
		return ErrInvalidOperationType
	}
	return nil
}

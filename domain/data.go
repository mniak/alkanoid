package domain

import "time"

type Account struct {
	ID             int
	DocumentNumber string
}

type OperationType int

type Transaction struct {
	ID            int
	AccountID     int
	OperationType OperationType
	Amount        float64
	EventDate     time.Time
}

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

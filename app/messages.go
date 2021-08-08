package app

/*
Para que as coisas fiquem mais simples, optei por
colocar informações de validação e de serialização
aqui mesmo.
*/

type (
	CreateAccountRequest struct {
		DocumentNumber string `json:"document_number" binding:"required"`
	}
	CreateAccountResponse struct {
		AccountID int `json:"account_id"`
	}
)

type (
	GetAccountRequest struct {
		AccountID int `uri:"accountId" binding:"required"`
	}
	GetAccountResponse struct {
		AccountID      int    `json:"account_id"`
		DocumentNumber string `json:"document_number"`
	}
)

type (
	CreateTransactionRequest struct {
		AccountID       int     `json:"account_id" binding:"required"`
		OperationTypeID int     `json:"operation_type_id" binding:"required"`
		Amount          float64 `json:"amount" binding:"required"`
	}
	CreateTransactionResponse struct {
		TransactionID int `json:"transaction_id"`
	}
)

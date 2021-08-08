package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/infra/persistence/inmemory"
)

func setupRouter(a app.Application) *gin.Engine {
	r := gin.Default()

	r.POST("/accounts", AppHandler(a, CreateAccount))
	r.GET("/accounts/:accountId", AppHandler(a, GetAccount))
	r.POST("/transactions", AppHandler(a, CreateTransaction))
	return r
}

func main() {
	a := app.NewApplication(
		app.RepositoriesRegistry{
			Account:     inmemory.NewAccountRepository(),
			Transaction: inmemory.NewTransactionRepository(),
		},
	)
	r := setupRouter(a)
	r.Run()
}

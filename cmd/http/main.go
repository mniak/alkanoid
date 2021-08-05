package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/application"
	"github.com/mniak/Alkanoid/infra/persistence/inmemory"
)

func setupRouter(app application.Application) *gin.Engine {
	r := gin.Default()

	r.POST("/accounts", AppHandler(app, CreateAccount))
	r.GET("/accounts/:accountId", AppHandler(app, GetAccount))
	r.POST("/transactions", AppHandler(app, CreateTransaction))
	return r
}

func main() {
	app := application.New(
		inmemory.NewAccountRepository(),
		inmemory.NewTransactionRepository(),
	)
	r := setupRouter(app)
	r.Run()
}

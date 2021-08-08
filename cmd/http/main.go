package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/infra/persistence/inmemory"
)

func setupRouter(a app.Application) *gin.Engine {
	router := gin.Default()
	router.POST("/accounts", AppHandler(a, CreateAccount))
	router.GET("/accounts/:accountId", AppHandler(a, GetAccount))
	router.POST("/transactions", AppHandler(a, CreateTransaction))
	return router
}

func main() {
	application := app.NewApplication(
		app.RepositoriesRegistry{
			Account:     inmemory.NewAccountRepository(),
			Transaction: inmemory.NewTransactionRepository(),
		},
	)
	engine := setupRouter(application)
	fmt.Println("Application started")
	engine.Run()
}

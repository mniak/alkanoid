package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/routes"
)

func main() {
	r := gin.Default()
	r.Group("")
	r.POST("/accounts", routes.POST_Accounts)
	r.GET("/accounts/:accountId", routes.GET_Accounts_ID)
	r.POST("/transactions", func(c *gin.Context) {
		c.Status(http.StatusCreated)
	})
	r.Run()
}

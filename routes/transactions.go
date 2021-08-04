package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(e *gin.Engine) {
	e.POST("/transactions", InsertTransaction)
}

func InsertTransaction(c *gin.Context) {
	c.Status(http.StatusCreated)
}

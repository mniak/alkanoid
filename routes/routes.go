package routes

import "github.com/gin-gonic/gin"

func RegisterAllRoutes(engine *gin.Engine) {
	RegisterTransactionRoutes(engine)
}

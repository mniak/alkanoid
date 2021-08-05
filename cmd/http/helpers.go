package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/app"
)

type AppHandlerFunc func(app.Application, *gin.Context) error

func AppHandler(a app.Application, fn AppHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(a, c)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/application"
)

type AppHandlerFunc func(application.Application, *gin.Context) error

func AppHandler(a application.Application, fn AppHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(a, c)
	}
}

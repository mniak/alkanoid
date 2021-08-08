package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/app"
	"github.com/mniak/Alkanoid/domain"
)

type AppHandlerFunc func(app.Application, *gin.Context) error

func AppHandler(a app.Application, fn AppHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := fn(a, c)
		if err == nil {
			return
		} else if domain.IsNotFoundError(err) {
			c.String(http.StatusNotFound, err.Error())
			return
		} else if domain.IsValidationError(err) {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.Status(http.StatusInternalServerError)
	}
}

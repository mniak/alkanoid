package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/application"
)

func CreateAccount(app application.Application, c *gin.Context) error {
	req := application.CreateAccountRequest{}
	_, err := app.CreateAccount(req)
	if err != nil {
		return err
	}
	c.Status(http.StatusCreated)
	return nil
}

func GetAccount(app application.Application, c *gin.Context) error {
	var req application.GetAccountRequest
	c.BindUri(&req)
	resp, err := app.GetAccount(req)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, resp)
	return nil
}

func CreateTransaction(app application.Application, c *gin.Context) error {
	req := application.CreateTransactionRequest{}
	_, err := app.CreateTransaction(req)
	if err != nil {
		return err
	}
	c.Status(http.StatusCreated)
	return nil
}

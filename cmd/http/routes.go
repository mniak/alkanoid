package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/application"
)

func CreateAccount(app application.Application, c *gin.Context) error {
	var req application.CreateAccountRequest
	c.BindJSON(&req)

	resp, err := app.CreateAccount(req)
	if err != nil {
		return err
	}
	c.JSON(http.StatusCreated, resp)
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
	var req application.CreateTransactionRequest
	c.BindJSON(&req)

	resp, err := app.CreateTransaction(req)
	if err != nil {
		return err
	}
	c.JSON(http.StatusCreated, resp)
	return nil
}

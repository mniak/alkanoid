package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/app"
)

func CreateAccount(a app.Application, c *gin.Context) error {
	var req app.CreateAccountRequest
	c.BindJSON(&req)

	resp, err := a.CreateAccount(req)
	if err != nil {
		return err
	}
	c.JSON(http.StatusCreated, resp)
	return nil
}

func GetAccount(a app.Application, c *gin.Context) error {
	var req app.GetAccountRequest
	c.BindUri(&req)

	resp, err := a.GetAccount(req)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, resp)
	return nil
}

func CreateTransaction(a app.Application, c *gin.Context) error {
	var req app.CreateTransactionRequest
	c.BindJSON(&req)

	resp, err := a.CreateTransaction(req)
	if err != nil {
		return err
	}
	c.JSON(http.StatusCreated, resp)
	return nil
}

package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/domain"
)

func GET_Accounts_ID(c *gin.Context) {
	accountIDString := c.Param("accountId")
	accountID, err := strconv.Atoi(accountIDString)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, domain.Account{
		ID:             accountID,
		DocumentNumber: "12345678900",
	})
}

func POST_Accounts(c *gin.Context) {
	c.Status(http.StatusCreated)
}

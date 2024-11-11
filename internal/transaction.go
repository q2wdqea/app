package internal

import (
	"app/pkg/model/response"
	"app/pkg/service"
	"github.com/gin-gonic/gin"
)

type transaction struct {
	Transaction *service.Transaction
	Wallet      *service.Wallet
}

func NewTransaction(t *service.Transaction, w *service.Wallet) *transaction {
	return &transaction{
		Transaction: t,
		Wallet:      w,
	}
}

func (t *transaction) Transfer(c *gin.Context) {
	response.Ok(c)
}

func (t *transaction) Deposit(c *gin.Context) {
	response.Ok(c)
}

func (t *transaction) Withdraw(c *gin.Context) {
	response.Ok(c)
}

func (t *transaction) ViewTransaction(c *gin.Context) {
	response.Ok(c)
}

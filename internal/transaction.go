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

// Transfer from one user to another user
func (t *transaction) Transfer(c *gin.Context) {
	response.Ok(c)
}

// Deposit to specify user wallet
func (t *transaction) Deposit(c *gin.Context) {
	response.Ok(c)
}

// Withdraw from specify user wallet
func (t *transaction) Withdraw(c *gin.Context) {
	response.Ok(c)
}

// ViewTransaction get specify user transaction history
func (t *transaction) ViewTransaction(c *gin.Context) {
	response.Ok(c)
}

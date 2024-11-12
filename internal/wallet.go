package internal

import (
	"app/pkg/model"
	"app/pkg/model/response"
	"app/pkg/service"
	"github.com/gin-gonic/gin"
)

type wallet struct {
	Transaction *service.Transaction
	Wallet      *service.Wallet
}

func NewWallet(t *service.Transaction, w *service.Wallet) *wallet {
	return &wallet{
		Transaction: t,
		Wallet:      w,
	}
}

// Transfer from one user to another user
func (w *wallet) Transfer(c *gin.Context) {
	var request *model.Transfer
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ParamError(c)
		return
	}
	if err := w.Wallet.Transfer(c, request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Deposit to specify user wallet
func (w *wallet) Deposit(c *gin.Context) {
	var request *model.Deposit
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ParamError(c)
		return
	}
	if err := w.Wallet.Deposit(c, request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Withdraw from specify user wallet
func (w *wallet) Withdraw(c *gin.Context) {
	var request *model.Withdraw
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ParamError(c)
		return
	}
	if err := w.Wallet.Withdraw(c, request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// ViewTransaction get specify user transaction history
func (w *wallet) ViewTransaction(c *gin.Context) {
	response.Ok(c)
}

// Balance Get specify user balance
func (w *wallet) Balance(c *gin.Context) {
	response.Ok(c)
}

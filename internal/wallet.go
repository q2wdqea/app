package internal

import (
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

// Balance Get specify user balance
func (w *wallet) Balance(c *gin.Context) {
	response.Ok(c)
}

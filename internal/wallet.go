package internal

import (
	"app/ecode"
	"app/pkg/model"
	"app/pkg/model/response"
	"app/pkg/service"
	"sort"
	"strconv"
	"sync"

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
	var request *model.TransferRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ParamError(c, err.Error())
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
	var request *model.DepositRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ParamError(c, err.Error())
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
	var request *model.WithdrawRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ParamError(c, err.Error())
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
	uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil || uid <= 0 {
		response.ParamError(c, ecode.Errors[ecode.ParamError])
		return
	}
	pageNo, _ := strconv.Atoi(c.Query("page_no"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize <= 0 {
		pageSize = 10
	}
	userId := int64(uid)
	transactions, err := w.Transaction.FindTransaction(c, userId, pageNo, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	total, err := w.Transaction.CountTransaction(c, userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	transferIds := make([]int64, 0)
	withdrawIds := make([]int64, 0)
	depositIds := make([]int64, 0)
	for _, v := range transactions {
		switch v.BizType {
		case model.TransactionTransfer.Int64():
			transferIds = append(transferIds, v.BizId)
		case model.TransactionWithdraw.Int64():
			withdrawIds = append(withdrawIds, v.BizId)
		case model.TransactionDeposit.Int64():
			depositIds = append(depositIds, v.BizId)
		}
	}
	var wg sync.WaitGroup
	async := &model.AsyncTransactions{
		Withdraws: make([]*model.TransactionDTO, 0),
		Deposits:  make([]*model.TransactionDTO, 0),
		Transfers: make([]*model.TransactionDTO, 0),
	}
	var list model.TransactionDTOS
	wg.Add(3)
	go w.Transaction.AsyncDeposits(c, depositIds, &wg, async)
	go w.Transaction.AsyncWithdraws(c, withdrawIds, &wg, async)
	go w.Transaction.AsyncTransfers(c, transferIds, &wg, async)
	wg.Wait()
	list = append(append(async.Deposits, async.Transfers...), async.Withdraws...)
	sort.Sort(list)
	response.OkWithData(&model.TransactionListDTO{Transactions: list, Total: total}, c)
}

// Balance Get specify user balance
func (w *wallet) Balance(c *gin.Context) {
	uid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil || uid <= 0 {
		response.ParamError(c, ecode.Errors[ecode.ParamError])
		return
	}
	balance, err := w.Wallet.FindOne(c, int64(uid))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(balance, c)
}

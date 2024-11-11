package main

import (
	"app/config"
	"app/internal"
	"app/pkg/db"
	"app/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// init db
	db := db.NewDB(config.Cfg)

	// init service
	walletService := service.NewWallet(db)
	transactionService := service.NewTransaction(db)

	// init controller
	walletController := internal.NewWallet(transactionService, walletService)
	transactionController := internal.NewTransaction(transactionService, walletService)

	// init web framework
	r := gin.Default()
	v1Group := r.Group("v1")
	{
		v1Group.GET("/balance", walletController.Balance)
		v1Group.GET("/view/transaction", transactionController.ViewTransaction)
		v1Group.GET("/deposit", transactionController.Deposit)
		v1Group.GET("/transfer", transactionController.Transfer)
		v1Group.GET("/withdraw", transactionController.Withdraw)
	}
	r.Run(":9090")
}

package main

import (
	"app/config"
	"app/internal"
	"app/pkg/cache"
	"app/pkg/db"
	"app/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// init db
	db := db.NewDB(config.Cfg)
	// init cache
	cache := cache.NewRedis(config.Cfg)

	// init service
	walletService := service.NewWallet(db, cache)
	transactionService := service.NewTransaction(db, cache)

	// init controller
	walletController := internal.NewWallet(transactionService, walletService)
	transactionController := internal.NewTransaction(transactionService, walletService)

	// init web framework
	r := gin.Default()
	v1Group := r.Group("v1")
	{
		v1Group.GET("/balance", walletController.Balance)
		v1Group.GET("/transactions", transactionController.ViewTransaction)
		v1Group.POST("/deposit", transactionController.Deposit)
		v1Group.POST("/transfer", transactionController.Transfer)
		v1Group.POST("/withdraw", transactionController.Withdraw)
	}
	r.Run(":9090")
}

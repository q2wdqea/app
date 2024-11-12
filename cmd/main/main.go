package main

import (
	"app/config"
	"app/internal"
	"app/pkg/cache"
	"app/pkg/db"
	"app/pkg/service"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {
	// init config
	config.InitConfig()
	ctx := context.Background()
	// new db
	db := db.NewDB(ctx, config.Cfg)
	defer db.Close()
	// new cache
	cache := cache.NewRedis(ctx, config.Cfg)

	// new service
	walletService := service.NewWallet(db, cache)
	transactionService := service.NewTransaction(db, cache)

	// new controller
	walletController := internal.NewWallet(transactionService, walletService)

	// set web framework api
	r := gin.Default()
	v1Group := r.Group("v1")
	{
		v1Group.GET("/balance", walletController.Balance)
		v1Group.GET("/transactions", walletController.ViewTransaction)
		v1Group.POST("/deposit", walletController.Deposit)
		v1Group.POST("/transfer", walletController.Transfer)
		v1Group.POST("/withdraw", walletController.Withdraw)
	}
	r.Run(":9090")
}

package service

import (
	"app/pkg/db"
	"context"
)

type wallet struct {
	DB *db.DB
}

func NewWallet(db *db.DB) *wallet {
	return &wallet{DB: db}
}

type WalletInterface interface {
	Withdraw(ctx context.Context,)
}

package service

import (
	"app/pkg/db"
	"app/pkg/model"
	"context"
)

type Wallet struct {
	DB *db.DB
}

func NewWallet(db *db.DB) *Wallet {
	return &Wallet{DB: db}
}

type WalletInterface interface {
	Withdraw(ctx context.Context, withdraw *model.Withdraw) error
	Transfer(ctx context.Context, withdraw *model.Transfer) error
	Balance(ctx context.Context, userId int64) (*model.Wallet, error)
	Deposit(ctx context.Context, deposit *model.Deposit) error
}

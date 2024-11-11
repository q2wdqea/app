package service

import (
	"app/pkg/cache"
	"app/pkg/db"
	"app/pkg/model"
	"context"
)

type Wallet struct {
	DB    *db.DB
	Cache *cache.Cache
}

func NewWallet(db *db.DB, cache *cache.Cache) *Wallet {
	return &Wallet{DB: db, Cache: cache}
}

type WalletInterface interface {
	Withdraw(ctx context.Context, withdraw *model.Withdraw) error
	Transfer(ctx context.Context, transfer *model.Transfer) error
	Balance(ctx context.Context, userId int64) (*model.Wallet, error)
	Deposit(ctx context.Context, deposit *model.Deposit) error
}

// Withdraw from specify user wallet
func (w *Wallet) Withdraw(ctx context.Context, withdraw *model.Withdraw) error {
	tx, err := w.DB.Begin()
	if err != nil {
		return err
	}
	status := true
	defer func() {
		if status {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "UPDATE t_wallet SET balance = balance - $1 WHERE user_id = $2 and balance > 0", withdraw.Amount, withdraw.UserId)
	if err != nil {
		status = false
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO t_withdraw (user_id, amount) VALUES ($1, $2)", withdraw.UserId, withdraw.Amount)
	if err != nil {
		status = false
		return err
	}
	return nil
}

// Transfer from one user to another user
func (w *Wallet) Transfer(ctx context.Context, transfer *model.Transfer) error {
	return nil
}

// Balance get specify user balance
func (w *Wallet) Balance(ctx context.Context, userId int64) (*model.Wallet, error) {
	var wallet *model.Wallet
	row := w.DB.QueryRowContext(ctx, "select id, user_id, balance, create_time from t_wallet where user_id = $1", userId)
	if err := row.Scan(&wallet); err != nil {
		return nil, err
	}
	return wallet, nil
}

// Deposit to specify user wallet
func (w *Wallet) Deposit(ctx context.Context, deposit *model.Deposit) error {
	tx, err := w.DB.Begin()
	if err != nil {
		return err
	}
	status := true
	defer func() {
		if status {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "UPDATE t_wallet SET balance = balance + $1 WHERE user_id = $2", deposit.Amount, deposit.UserId)
	if err != nil {
		status = false
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO t_deposit (user_id, amount) VALUES ($1, $2)", deposit.UserId, deposit.Amount)
	if err != nil {
		status = false
		return err
	}
	return nil
}

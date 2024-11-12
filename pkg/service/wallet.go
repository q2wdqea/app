package service

import (
	"app/pkg/cache"
	"app/pkg/db"
	"app/pkg/model"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Wallet struct {
	DB    *db.DB
	Cache *cache.Cache
}

func NewWallet(db *db.DB, cache *cache.Cache) *Wallet {
	return &Wallet{DB: db, Cache: cache}
}

type WalletInterface interface {
	Withdraw(ctx context.Context, withdraw *model.WithdrawRequest) error
	Transfer(ctx context.Context, transfer *model.TransferRequest) error
	FindOne(ctx context.Context, userId int64) (*model.Wallet, error)
	Deposit(ctx context.Context, deposit *model.DepositRequest) error
}

// Withdraw from specify user wallet
func (w *Wallet) Withdraw(ctx context.Context, withdraw *model.WithdrawRequest) error {
	if withdraw.Amount <= 0 {
		return errors.New("the withdrawal amount must be greater than 0")
	}
	wallet, err := w.FindOne(ctx, withdraw.UserId)
	if err != nil {
		return fmt.Errorf("wallet record not found user_id:%d", withdraw.UserId)
	}
	if wallet.Balance < withdraw.Amount {
		return errors.New("insufficient wallet balance")
	}
	toId := strconv.Itoa(int(withdraw.UserId))
	keys := []string{toId}
	defer func() {
		w.Cache.ReleaseMultipleLock(ctx, keys)
	}()
	lock := w.Cache.LockMultipleKeys(ctx, keys, 30*time.Second)
	if !lock {
		return errors.New("withdraw lock multiple keys lock failed to id: " + toId)
	}
	tx, err := w.DB.Begin()
	if err != nil {
		return err
	}
	rollback := true
	defer func() {
		if rollback {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	rst, err := tx.ExecContext(ctx, "UPDATE t_wallet SET balance = balance - $1 WHERE user_id = $2 and balance - $3 >= 0", withdraw.Amount, withdraw.UserId, withdraw.Amount)
	if err != nil {
		return err
	}
	if affected, _ := rst.RowsAffected(); affected < 1 {
		return errors.New("withdraw rows affected balance less than 0")
	}
	var lastInsertID int
	if err = tx.QueryRowContext(ctx, "INSERT INTO t_withdraw (user_id, amount) VALUES ($1, $2) RETURNING id", withdraw.UserId, withdraw.Amount).Scan(&lastInsertID); err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO t_transaction (user_id, biz_type, biz_id) VALUES ($1, $2, $3)",
		withdraw.UserId, model.TransactionWithdraw, lastInsertID)
	if err != nil {
		return err
	}
	rollback = false
	return nil
}

// Transfer from one user to another user
func (w *Wallet) Transfer(ctx context.Context, transfer *model.TransferRequest) error {
	if transfer.Amount <= 0 {
		return errors.New("the transfer amount must be greater than 0")
	}
	wallet, err := w.FindOne(ctx, transfer.FromId)
	if err != nil {
		return fmt.Errorf("send wallet record not found user_id:%d", transfer.FromId)
	}
	if wallet.Balance < transfer.Amount {
		return errors.New("insufficient wallet balance")
	}
	if _, err = w.FindOne(ctx, transfer.ToId); err != nil {
		return fmt.Errorf("receive wallet record not found user_id:%d", transfer.ToId)
	}
	fromId := strconv.Itoa(int(transfer.FromId))
	toId := strconv.Itoa(int(transfer.ToId))
	keys := []string{fromId, toId}
	defer func() {
		w.Cache.ReleaseMultipleLock(ctx, keys)
	}()
	lock := w.Cache.LockMultipleKeys(ctx, keys, 30*time.Second)
	if !lock {
		return errors.New("transfer lock multiple keys lock failed from id: " + fromId + " to id:" + toId)
	}
	tx, err := w.DB.Begin()
	if err != nil {
		return err
	}
	rollback := true
	defer func() {
		if rollback {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.ExecContext(ctx, "UPDATE t_wallet SET balance = balance + $1 WHERE user_id = $2", transfer.Amount, transfer.ToId)
	if err != nil {
		return err
	}
	rst, err := tx.ExecContext(ctx, "UPDATE t_wallet SET balance = balance - $1 WHERE user_id = $2 and balance - $3 >= 0", transfer.Amount, transfer.FromId, transfer.Amount)
	if err != nil {
		return err
	}
	if affected, _ := rst.RowsAffected(); affected < 1 {
		return errors.New("transfer rows affected balance less than 0")
	}
	var lastInsertID int
	if err = tx.QueryRowContext(ctx, "INSERT INTO t_transfer (from_id, to_id, amount) VALUES ($1, $2, $3) RETURNING id",
		transfer.FromId, transfer.ToId, transfer.Amount).Scan(&lastInsertID); err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO t_transaction (user_id, biz_type, biz_id) VALUES ($1, $2, $3),($4, $5, $6)",
		transfer.FromId, model.TransactionTransfer, lastInsertID, transfer.ToId, model.TransactionTransfer, lastInsertID)
	if err != nil {
		return err
	}
	rollback = false
	return nil
}

// FindOne get specify user balance
func (w *Wallet) FindOne(ctx context.Context, userId int64) (*model.Wallet, error) {
	var wallet model.Wallet
	row := w.DB.QueryRowContext(ctx, "select id, user_id, balance, create_time from t_wallet where user_id = $1", userId)
	if err := row.Scan(&wallet.Id, &wallet.UserId, &wallet.Balance, &wallet.CreateTime); err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Deposit to specify user wallet
func (w *Wallet) Deposit(ctx context.Context, deposit *model.DepositRequest) error {
	if deposit.Amount <= 0 {
		return errors.New("the deposit amount must be greater than 0")
	}
	if _, err := w.FindOne(ctx, deposit.UserId); err != nil {
		return fmt.Errorf("wallet record not found user_id:%d", deposit.UserId)
	}
	fromId := strconv.Itoa(int(deposit.UserId))
	keys := []string{fromId}
	defer func() {
		w.Cache.ReleaseMultipleLock(ctx, keys)
	}()
	lock := w.Cache.LockMultipleKeys(ctx, keys, 30*time.Second)
	if !lock {
		return errors.New("deposit lock multiple keys lock failed from id: " + fromId)
	}
	tx, err := w.DB.Begin()
	if err != nil {
		return err
	}
	rollback := true
	defer func() {
		if rollback {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.ExecContext(ctx, "UPDATE t_wallet SET balance = balance + $1 WHERE user_id = $2", deposit.Amount, deposit.UserId)
	if err != nil {
		return err
	}
	var lastInsertID int
	if err = tx.QueryRowContext(ctx, "INSERT INTO t_deposit (user_id, amount) VALUES ($1, $2) RETURNING id", deposit.UserId, deposit.Amount).Scan(&lastInsertID); err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO t_transaction (user_id, biz_type, biz_id) VALUES ($1, $2, $3)",
		deposit.UserId, model.TransactionDeposit, lastInsertID)
	if err != nil {
		return err
	}
	rollback = false
	return nil
}

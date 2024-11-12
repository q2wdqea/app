package service

import (
	"app/pkg/cache"
	"app/pkg/db"
	"app/pkg/model"
	"context"
)

type Transaction struct {
	DB    *db.DB
	Cache *cache.Cache
}

func NewTransaction(db *db.DB, cache *cache.Cache) *Transaction {
	return &Transaction{DB: db, Cache: cache}
}

type TransactionInterface interface {
	FindTransaction(ctx context.Context, userId int64, page, size int64) ([]*model.Transaction, error)
	CountTransaction(ctx context.Context, userId int64) (int64, error)
}

// FindTransaction find specify user transaction history
func (t *Transaction) FindTransaction(ctx context.Context, userId int64, page, size int64) ([]*model.Transaction, error) {
	transactions := make([]*model.Transaction, 0)
	rows, err := t.DB.QueryContext(ctx, "select id, from_id, to_id, amount, create_time from t_transfer where from_id = $1 or to_id = $2", userId, userId)
	if err != nil {
		return nil, err
	}
	if err := rows.Scan(&transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}

// CountTransaction count specify user transaction history
func (t *Transaction) CountTransaction(ctx context.Context, userId int64) (int64, error) {
	var transactions int64
	row := t.DB.QueryRowContext(ctx, "select count(1) from t_transfer where from_id = $1 or to_id = $2", userId, userId)
	if err := row.Scan(&transactions); err != nil {
		return 0, err
	}
	return transactions, nil
}

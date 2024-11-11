package service

import (
	"app/pkg/db"
	"app/pkg/model"
	"context"
)

type transaction struct {
	DB *db.DB
}

func NewTransaction(db *db.DB) *transaction {
	return &transaction{DB: db}
}

type TransactionInterface interface {
	FindTransaction(ctx context.Context, userId int64) ([]*model.Transaction, error)
	CountTransaction(ctx context.Context, userId int64) (int64, error)
}

func (t *transaction) FindTransaction(ctx context.Context, userId int64) ([]*model.Transaction, error) {
	return nil, nil
}

func (t *transaction) CountTransaction(ctx context.Context, userId int64) (int64, error) {
	return 0, nil
}

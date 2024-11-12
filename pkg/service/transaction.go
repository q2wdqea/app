package service

import (
	"app/pkg/cache"
	"app/pkg/db"
	"app/pkg/model"
	"context"
	"log"
	"sync"

	"github.com/lib/pq"
)

type Transaction struct {
	DB    *db.DB
	Cache *cache.Cache
}

func NewTransaction(db *db.DB, cache *cache.Cache) *Transaction {
	return &Transaction{DB: db, Cache: cache}
}

type TransactionInterface interface {
	FindTransaction(ctx context.Context, userId int64, page, size int) ([]*model.Transaction, error)
	CountTransaction(ctx context.Context, userId int64) (int64, error)
	AsyncTransfers(ctx context.Context, bizIds []int64, wg *sync.WaitGroup, async *model.AsyncTransactions)
	AsyncWithdraws(ctx context.Context, bizIds []int64, wg *sync.WaitGroup, async *model.AsyncTransactions)
	AsyncDeposits(ctx context.Context, bizIds []int64, wg *sync.WaitGroup, async *model.AsyncTransactions)
}

// AsyncTransfers async transfer user transaction history
func (t *Transaction) AsyncTransfers(ctx context.Context, bizIds []int64, wg *sync.WaitGroup, async *model.AsyncTransactions) {
	defer wg.Done()
	if len(bizIds) == 0 {
		return
	}
	query := "select id, from_id, to_id, amount, create_time from t_transfer where id = ANY($1)"
	rows, err := t.DB.QueryContext(ctx, query, pq.Array(bizIds))
	if err != nil {
		log.Fatalf("async transfer err: %v", err)
		return
	}
	for rows.Next() {
		var transfer model.TransactionDTO
		if err := rows.Scan(&transfer.Id, &transfer.FromId, &transfer.ToId, &transfer.Amount, &transfer.CreateTime); err != nil {
			log.Fatalf("db scan err: %v", err)
		}
		transfer.BizType = model.TransactionTransfer.String()
		async.Transfers = append(async.Transfers, &transfer)
	}
}

// AsyncDeposits async deposit user transaction history
func (t *Transaction) AsyncDeposits(ctx context.Context, bizIds []int64, wg *sync.WaitGroup, async *model.AsyncTransactions) {
	defer wg.Done()
	if len(bizIds) == 0 {
		return
	}
	query := "select id, user_id,  amount, create_time from t_deposit where id = ANY($1)"
	rows, err := t.DB.QueryContext(ctx, query, pq.Array(bizIds))
	if err != nil {
		log.Fatalf("async deposit err: %v", err)
		return
	}
	for rows.Next() {
		var deposit model.TransactionDTO
		if err := rows.Scan(&deposit.Id, &deposit.UserId, &deposit.Amount, &deposit.CreateTime); err != nil {
			log.Fatalf("db scan err: %v", err)
			return
		}
		deposit.BizType = model.TransactionDeposit.String()
		async.Deposits = append(async.Deposits, &deposit)
	}
}

// AsyncWithdraws async withdraw user transaction history
func (t *Transaction) AsyncWithdraws(ctx context.Context, bizIds []int64, wg *sync.WaitGroup, async *model.AsyncTransactions) {
	defer wg.Done()
	if len(bizIds) == 0 {
		return
	}
	query := "select id, user_id,  amount, create_time from t_withdraw where id = ANY($1)"
	rows, err := t.DB.QueryContext(ctx, query, pq.Array(bizIds))
	if err != nil {
		log.Fatalf("async withdraw err: %v", err)
		return
	}
	for rows.Next() {
		var withdraw model.TransactionDTO
		if err := rows.Scan(&withdraw.Id, &withdraw.UserId, &withdraw.Amount, &withdraw.CreateTime); err != nil {
			log.Fatalf("db scan err: %v", err)
			return
		}
		withdraw.BizType = model.TransactionWithdraw.String()
		async.Withdraws = append(async.Withdraws, &withdraw)
	}
}

// FindTransaction find specify user transaction history
func (t *Transaction) FindTransaction(ctx context.Context, userId int64, page, size int) ([]*model.Transaction, error) {
	transactions := make([]*model.Transaction, 0)
	offset := (page - 1) * size
	if offset < 0 {
		offset = 0
	}
	if size < 0 {
		size = 10
	}
	rows, err := t.DB.QueryContext(ctx, "select id, user_id, biz_type, biz_id, create_time from t_transaction where user_id = $1 limit $2 offset $3", userId, size, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.UserId, &transaction.BizType, &transaction.BizId, &transaction.CreateTime); err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}

// CountTransaction count specify user transaction history
func (t *Transaction) CountTransaction(ctx context.Context, userId int64) (int64, error) {
	var transactions int64
	row := t.DB.QueryRowContext(ctx, "select count(1) from t_transaction where user_id = $1", userId)
	if err := row.Scan(&transactions); err != nil {
		return 0, err
	}
	return transactions, nil
}

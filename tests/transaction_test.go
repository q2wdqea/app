package tests

import (
	"app/pkg/cache"
	"app/pkg/db"
	"app/pkg/model"
	"app/pkg/service"
	"context"
	"database/sql"
	"fmt"
	"sort"
	"sync"
	"testing"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

const (
	host   = "127.0.0.1"
	user   = "postgres"
	pwd    = "hs2024!@#"
	dbname = "app"
	addr   = "127.0.0.1:6379"
)

func NewDBAndCache(t *testing.T) (*db.DB, *cache.Cache) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	if _, err := client.Ping(ctx).Result(); err != nil {
		t.Errorf("redis connect err: %v", err)
	}
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", user, pwd, host, dbname)
	dataSource, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Errorf("mysql connect err: %v", err)
	}
	if err = dataSource.PingContext(ctx); err != nil {
		t.Errorf("mysql connect err: %v", err)
	}
	return &db.DB{DB: dataSource}, &cache.Cache{Client: client}
}

func TestFindTransaction(t *testing.T) {
	ctx := context.Background()
	db, cache := NewDBAndCache(t)
	transactionService := service.NewTransaction(db, cache)
	list, err := transactionService.FindTransaction(ctx, 1, 1, 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(list)

	list2, err := transactionService.FindTransaction(ctx, 1, -1, 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(list2)

	list3, err := transactionService.FindTransaction(ctx, 1, -1, -10)
	if err != nil {
		t.Error(err)
	}
	t.Log(list3)

	list4, err := transactionService.FindTransaction(ctx, -1, 1, 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(list4)

	list5, err := transactionService.FindTransaction(ctx, -1, -1, 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(list5)

	list6, err := transactionService.FindTransaction(ctx, -1, 1, -10)
	if err != nil {
		t.Error(err)
	}
	t.Log(list6)
}

func TestCountTransaction(t *testing.T) {
	ctx := context.Background()
	db, cache := NewDBAndCache(t)
	transactionService := service.NewTransaction(db, cache)
	c1, err := transactionService.CountTransaction(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(c1)

	c2, err := transactionService.CountTransaction(ctx, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(c2)

	c3, err := transactionService.CountTransaction(ctx, -1)
	if err != nil {
		t.Error(err)
	}
	t.Log(c3)

	c4, err := transactionService.CountTransaction(ctx, 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(c4)
}

func TestAsyncTransactions(t *testing.T) {
	ctx := context.Background()
	db, cache := NewDBAndCache(t)
	transactionService := service.NewTransaction(db, cache)
	ts, err := transactionService.FindTransaction(ctx, 1, 1, 10)
	if err != nil {
		t.Error(err)
	}
	t.Log(ts)
	transferIds := make([]int64, 0)
	withdrawIds := make([]int64, 0)
	depositIds := make([]int64, 0)
	for _, v := range ts {
		switch v.BizType {
		case model.TransactionTransfer.Int64():
			transferIds = append(transferIds, v.BizId)
		case model.TransactionWithdraw.Int64():
			withdrawIds = append(withdrawIds, v.BizId)
		case model.TransactionDeposit.Int64():
			depositIds = append(depositIds, v.BizId)
		}
	}
	var wg sync.WaitGroup
	async := &model.AsyncTransactions{
		Withdraws: make([]*model.TransactionDTO, 0),
		Deposits:  make([]*model.TransactionDTO, 0),
		Transfers: make([]*model.TransactionDTO, 0),
	}
	var list model.TransactionDTOS
	wg.Add(3)
	go transactionService.AsyncDeposits(ctx, depositIds, &wg, async)
	go transactionService.AsyncWithdraws(ctx, withdrawIds, &wg, async)
	go transactionService.AsyncTransfers(ctx, transferIds, &wg, async)
	wg.Wait()
	list = append(append(async.Deposits, async.Transfers...), async.Withdraws...)
	sort.Sort(list)
	t.Logf("list data %v:", list)
}

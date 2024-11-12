package db

import (
	"app/config"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB(ctx context.Context, cfg *config.Config) *DB {
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}
	return &DB{DB: db}
}

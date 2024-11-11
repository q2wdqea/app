package db

import (
	"app/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB(cfg *config.Config) *DB {
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	return &DB{DB: db}
}

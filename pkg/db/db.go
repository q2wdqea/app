package db

import (
	"database/sql"
	"fmt"
)

type DB struct {
	DB *sql.DB
}

func NewDB(uname, pwd, dbname string) *DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", uname, pwd, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	return &DB{DB: db}
}

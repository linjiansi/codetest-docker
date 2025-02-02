package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root@tcp(db:3306)/codetest")
	if err != nil {
		return nil, err
	}

	return db, nil
}

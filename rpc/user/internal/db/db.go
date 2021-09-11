package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Init(mysqlDSN string) {
	db = sqlx.MustConnect("mysql", mysqlDSN)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(3)
}

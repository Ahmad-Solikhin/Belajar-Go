package Goalng_Database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}

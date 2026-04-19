package Goalng_Database

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

package Goalng_Database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "insert into customer(id, name) values('gayuh', 'Gayuh')")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "select id, name from customer")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}

	fmt.Println("Success query new customer")
}

type Customer struct {
	Id        string
	Name      string
	Balance   int64
	Rating    sql.NullFloat64
	CreatedAt time.Time
	BirthDate time.Time
	Married   bool
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "select id, name, balance, rating, created_at, birth_date, married from customer")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Balance, &customer.Rating, &customer.CreatedAt, &customer.BirthDate, &customer.Married)
		if err != nil {
			panic(err)
		}
		fmt.Println("Customer : ", customer)
	}

	fmt.Println("Success query new customer")
}

func TestAutoIncrementSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	var id int64

	script := "insert into comment(comment) values($1) returning id"
	err := db.QueryRowContext(ctx, script, "Haiyaaa Looo ke-2").Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Last insert id ", id)

	fmt.Println("Success insert new comment")
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	stmt, err := db.PrepareContext(ctx, "insert into comment(comment) values($1) returning id")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var id int64
	err = stmt.QueryRowContext(ctx, "Haiyaaa pake prepared statement").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Last insert id ", id)
	fmt.Println("Success insert new comment")
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	script := "insert into comment(comment) values($1)"
	_, err = tx.ExecContext(ctx, script, "Pake prepared statement")
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}

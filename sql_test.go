package Belajar_Database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {

		panic(err)
	}
	fmt.Println("Berhasil")
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
}

func TestQuerySql(t *testing.T) {
	db := getConnection()
	defer db.Close()
	ctx := context.Background()
	sql := "Select id, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id, name, email      string
			balance              int32
			rating               float64
			createdAt, birthDate time.Time
			married              bool
		)
		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name, email, balance, rating, createdAt, birthDate, married)
	}
}

func TestSqlInjection(t *testing.T) {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #"
	password := "ha"

	script := "SELECT username FROM user WHERE username = '" + username + "'AND password = '" + password + "'LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {

		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses")
	} else {
		fmt.Println("Gagal")
	}
}

func TestSqlInjectionFix(t *testing.T) {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #"
	password := "ha"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {

		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses")
	} else {
		fmt.Println("Gagal")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()
	email := "Bagus@gmail.com"
	comment := "Test apakah berhasil"

	script := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

func TestPrepareStatement(t *testing.T) {
	db := getConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini comment ke-" + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, _ := result.LastInsertId()
		fmt.Println(id)
	}

}

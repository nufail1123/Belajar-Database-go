package Belajar_Database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpthy(t *testing.T) {}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_database_golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

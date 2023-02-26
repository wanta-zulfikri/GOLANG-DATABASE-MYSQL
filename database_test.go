package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// Open Connexion
func TestOpenConnection(t *testing.T){

db, err := sql.Open("mysql", "root:golang@tcp(host:3306)/belajar_golang_database")
if err != nil {
	panic (err)
}

defer db.Close()
// Gunakan DB 
}


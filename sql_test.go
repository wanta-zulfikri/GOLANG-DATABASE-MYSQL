package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)
// menambahkan data 
func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(Id, name) VALUE('eko, 'eko)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("succes insert new customer")

}

// mengelolah data 
func TestQuerySql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name from customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Id:", )
	fmt.Println("Name:",)

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic (err)
		}


	}
}



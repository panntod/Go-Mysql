package main

import (
	"context"
	"fmt"
	"testing"
)

func TestExecInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	insertSql := "INSERT INTO `go-user` (nama, umur, alamat) VALUES (?, ?, ?)"
	_, err := db.ExecContext(ctx, insertSql, "dharma", 12, "malang")
	if err != nil {
		panic(err)
	}

	fmt.Println("Data Successfully Inserted")
}

func TestExecQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	selectSql := "SELECT * FROM `go-user`"
	row, err := db.QueryContext(ctx, selectSql)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		var (
			id     int
			nama   string
			umur   int
			alamat string
		)

		err = row.Scan(&id, &nama, &umur, &alamat)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id: %d, nama: %s, umur: %d, alamat: %s \n", id, nama, umur, alamat)
	}
}

func TestExecUpdateSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	updateSql := "UPDATE `go-user` SET umur = ? WHERE id = ?"
	_, err := db.ExecContext(ctx, updateSql, 20, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data Successfully Updated")
}

func TestExecDeleteSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	deleteSql := "DELETE FROM `go-user` WHERE id = ?"
	_, err := db.ExecContext(ctx, deleteSql, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data Successfully Deleted")
}

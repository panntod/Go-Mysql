package gomysql

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	insertSql := "INSERT INTO `go-user` (nama, umur, alamat) VALUES ('asfina', 16, 'kepanjen')"
	_, err := db.ExecContext(ctx, insertSql)
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

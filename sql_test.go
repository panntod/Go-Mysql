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

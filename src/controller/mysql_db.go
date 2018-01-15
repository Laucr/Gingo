package controller

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func connectMysql(dataSN string) (*sql.DB, int) {
	db, err := sql.Open("mysql", dataSN)
	if err != nil {
		fmt.Println("error", err)
		return nil, ConnectErr
	}
	return db, OperationSuccess
}

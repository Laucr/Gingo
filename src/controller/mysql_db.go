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

func InsertUserBasicInfo(u *UserBasicInfo) int {
	db, e := connectMysql(MysqlDB)
	if db == nil {
		fmt.Println("Error Code", e)
		return ConnectErr
	}
	insert := "INSERT INTO `users`.`basic_info` " +
		"(`UserId`, `Password`, `UserName`, `Email`, `Tel`, `CreateTime`) " +
		"VALUES (?,?,?,?,?,?);"
	ret, err := db.Exec(insert, u.UserId, u.Password, u.UserName, u.Email, u.Tel, u.CreateTime)
	if ret == nil {
		fmt.Println("Error", err)
		return InsertFailed
	}
	defer db.Close()
	return InsertSuccess
}

func SelectUserInfo(s string, k string) int {
	db, e := connectMysql(MysqlDB)
	if db == nil {
		fmt.Println("Error Code", e)
		return ConnectErr
	}
	return OperationSuccess
}
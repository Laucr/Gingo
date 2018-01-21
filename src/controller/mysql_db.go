package controller

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strings"
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
	return InsertSuccess
}

func SQLiProof(str string) int {
	if strings.ContainsAny(str, "`' %&*$#") {
		return -1
	}
	return 0
}

func SelectUserInfo(key string, value string) int {
	db, e := connectMysql(MysqlDB)
	if db == nil {
		fmt.Println("Error Code", e)
		return ConnectErr
	}
	searchUser := "SELECT UserId FROM `users`.`basic_info` WHERE " + key + "==" + value
	row, err := db.Query(searchUser)
	return OperationSuccess
}
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

func InsertUserBasicInfo(u *UserBasicInfo, password string) int {
	db, e := connectMysql(MysqlDB)
	if db == nil {
		fmt.Println("Error Code", e)
		return ConnectErr
	}
	insert := "INSERT INTO `users`.`basic_info` " +
		"(`UserId`, `Password`, `UserName`, `Email`, `Tel`, `CreateTime`) " +
		"VALUES (?,?,?,?,?,?);"
	ret, err := db.Exec(insert, u.UserId, password, u.UserName, u.Email, u.Tel, u.CreateTime)
	if ret == nil {
		fmt.Println("Error", err)
		return InsertFailed
	}
	return InsertSuccess
}

func SQLiProof(str string) int {
	if strings.ContainsAny(str, "`' %&*$#") {
		return SQLiParameter
	}
	return OperationSuccess
}

func GetUid(key string, value string) (int, int) {
	db, e := connectMysql(MysqlDB)
	uid := 0
	if SQLiProof(value) != OperationSuccess {
		return uid, SQLiParameter
	}

	if db == nil {
		fmt.Println("Error Code", e)
		return uid, ConnectErr
	}

	var query string

	switch key {
	//case "UserName":
	//	query = "SELECT UserId from `users`.`basic_info` WHERE `UserName` = ?"
	case "Email":
		query = "SELECT UserId from `users`.`basic_info` WHERE `Email` = ?"
	case "Tel":
		query = "SELECT UserId from `users`.`basic_info` WHERE `Tel` = ?"
	}
	row := db.QueryRow(query, value)

	if row == nil {
		return uid, QueryRowNotExists
	}

	if err := row.Scan(&uid); err != nil {
		return uid, QueryFailed
	}

	return uid, OperationSuccess
}

func UpdateUserBasicInfo(uid int, key string, value string) int {
	db, e := connectMysql(MysqlDB)
	if SQLiProof(value) != OperationSuccess {
		return SQLiParameter
	}

	if db == nil {
		fmt.Println("Error Code", e)
		return ConnectErr
	}

	tx, err := db.Begin()
	if tx == nil {
		fmt.Println("Error", err)
		return OperationFailed
	}
	var update string
	switch key {
	case "UserName":
		update = "UPDATE `users`.`basic_info` SET `UserName` = ? WHERE `UserId` = ?"
	case "Email":
		update = "UPDATE `users`.`basic_info` SET `Email` = ? WHERE `UserId` = ?"
	case "Tel":
		update = "UPDATE `users`.`basic_info` SET `Tel` = ? WHERE `UserId` = ?"
	}

	if _, err = tx.Exec(update, value, uid); err != nil {
		fmt.Println("Error", err)
		return OperationFailed
	}
	if err = tx.Commit(); err != nil {
		fmt.Println("Error", err)
		return OperationFailed
	}
	return OperationSuccess
}

func UpdatePassword(old string, new string, uid int) int {
	db, e := connectMysql(MysqlDB)
	if db == nil {
		fmt.Println("Error Code", e)
		return ConnectErr
	}

	tx, err := db.Begin()
	if tx == nil {
		fmt.Println("Error", err)
		return OperationFailed
	}

	query := "SELECT Password from `users`.`basic_info` WHERE `UserId` = ?"
	row := tx.QueryRow(query, uid)
	if row == nil {
		return QueryRowNotExists
	}

	var password string
	if err := row.Scan(&password); err != nil {
		return QueryFailed
	}
	if old != password {
		return PasswordInvalid
	}

	update := "UPDATE `users`.`basic_info` SET `Password` = ? WHERE `UserId` = ?"
	if _, err := tx.Exec(update, new, uid); err != nil {
		fmt.Println("Error", err)
		return OperationFailed
	}
	defer tx.Commit()
	return OperationSuccess
}

func CheckPassword(uid int, password string) int {
	db, e := connectMysql(MysqlDB)
	if db == nil {
		fmt.Println("Error Code", e)
		return ConnectErr
	}

	query := "SELECT Password FROM `users`.`basic_info` WHERE `UserId` = ?"
	row := db.QueryRow(query, uid)

	if row == nil {
		return QueryRowNotExists
	}

	var userPassword string
	if err := row.Scan(&userPassword); err != nil {
		return QueryFailed
	}
	if userPassword == password {
		return PasswordInvalid
	}

	return PasswordCorrect
}

func GetUserInfo(uid int) (*UserBasicInfo, *UserAdvInfo, int) {
	db, e := connectMysql(MysqlDB)

	if db == nil {
		fmt.Println("Error Code", e)
		return nil, nil, ConnectErr
	}
	//get basic user info
	queryBasicInfo := "SELECT `Email`, `Tel`, `UserName`, `CreateTime` FROM `user`.`basic_info` WHERE `UserId` = ?"
	var email, tel, username string
	var ct int
	row := db.QueryRow(queryBasicInfo, uid)
	if row == nil {
		return nil, nil, QueryFailed
	}
	row.Scan(&email, &tel, &username, &ct)
	basicInfo := UserBasicInfo{uid, email, tel, username, ct}

	var city, bars string
	queryAdvInfo := "SELECT `City`, `Bars` FROM `users`.`advance_info` WHERE `UserId` = ?"
	row = db.QueryRow(queryAdvInfo, uid)
	if row == nil {
		return &basicInfo, new(UserAdvInfo), LackOfAdvInfo
	}
	row.Scan(&city, &bars)
	advInfo := UserAdvInfo{city, bars}
	//get advance user info
	return &basicInfo, &advInfo, OperationSuccess
}

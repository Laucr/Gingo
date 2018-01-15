package controller

import (
	"strconv"
	"encoding/json"
)

type Users struct {
	userId   int64
	userName string
	password string
}

type UserBasicInfo struct {
	UserId     int
	Password   string
	UserName   string
	Email      string
	Tel        string
	CreateTime int
}

type UserAdvInfo struct {
	UserId int
	City string
	Bars string
}

type JsonClosure struct {
	js string
}

func JsonConvertToMap(j *JsonClosure) (map[string]interface{}, int) {
	var ret map[string]interface{}
	if err := json.Unmarshal([]byte(j.js), &ret); err != nil {
		return nil, OperationFailed
	}
	return ret, OperationSuccess
}

func UsersConvertToMap(user *Users) map[string]interface{} {
	fields := make(map[string]interface{})
	fields["userName"] = user.userName
	fields["userId"] = user.userId
	// TODO: encrypt password here
	fields["password"] = user.password
	return fields
}

func MapConvertToUser(fields map[string]string) *Users {
	u := new(Users)
	u.userName = fields["userName"]
	u.password = fields["password"]
	u.userId, _ = strconv.ParseInt(fields["userId"], 10, 64)
	return u
}

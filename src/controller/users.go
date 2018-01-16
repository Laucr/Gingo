package controller

import (
	"strconv"
	"encoding/json"
)

type Users struct {
	UserId   int64
	UserName string
	Password string
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

func ObjConvertToJson(v interface{}) string {
	var j string
	b, _ := json.Marshal(v)
	j = string(b)
	return j
}

func UsersConvertToMap(user *Users) map[string]interface{} {
	fields := make(map[string]interface{})
	fields["UserName"] = user.UserName
	fields["UserId"] = user.UserId
	// TODO: encrypt Password here
	fields["Password"] = user.Password
	return fields
}

func MapConvertToUser(fields map[string]string) *Users {
	u := new(Users)
	u.UserName = fields["UserName"]
	u.Password = fields["Password"]
	u.UserId, _ = strconv.ParseInt(fields["UserId"], 10, 64)
	return u
}

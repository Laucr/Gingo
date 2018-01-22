package controller

import (
	"encoding/json"
)

type Users struct {
	UserId    int            `json:"user_id"`
	Password  string         `json:"password"`
	BasicInfo *UserBasicInfo `json:"basic_info"`
	AdvInfo   *UserAdvInfo   `json:"adv_info"`
}

type UserBasicInfo struct {
	UserId     int    `json:"user_id"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
	UserName   string `json:"user_name"`
	CreateTime int    `json:"create_time"`
}

type UserAdvInfo struct {
	City string `json:"city"`
	Bars string `json:"bars"`
}

type JsonClosure struct {
	js string
}

//func JsonConvertToMap(j *JsonClosure) (map[string]interface{}, int) {
//	var ret map[string]interface{}
//	if err := json.Unmarshal([]byte(j.js), &ret); err != nil {
//		return nil, OperationFailed
//	}
//	return ret, OperationSuccess
//}

func ObjConvertToJson(v interface{}) string {
	var j string
	b, _ := json.Marshal(v)
	j = string(b)
	return j
}

//func UsersConvertToMap(user *Users) map[string]interface{} {
//	fields := make(map[string]interface{})
//	fields["UserName"] = user.UserName
//	fields["UserId"] = user.UserId
//	// TODO: encrypt Password here
//	fields["Password"] = user.Password
//	return fields
//}

//func MapConvertToUser(fields map[string]string) *Users {
//	u := new(Users)
//	u.UserName = fields["UserName"]
//	u.Password = fields["Password"]
//	u.UserId, _ = strconv.ParseInt(fields["UserId"], 10, 64)
//	return u
//}

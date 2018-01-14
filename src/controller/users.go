package controller

import "strconv"

type Users struct {
	userName string
	password string
	userId   int64
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

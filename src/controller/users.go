package controller

import "strconv"

type Users struct {
	username string
	email    string
	userId   int64
	password string
}

func UsersConvertToMap(user *Users) map[string]interface{} {
	fields := make(map[string]interface{})
	fields["username"] = user.username
	fields["email"] = user.email
	fields["userId"] = user.userId
	// TODO: encrypt password here
	fields["password"] = user.password
	return fields
}

func MapConvertToUser(fields map[string]string) *Users {
	u := new(Users)
	u.username = fields["username"]
	u.password = fields["password"]
	u.email = fields["email"]
	u.userId, _ = strconv.ParseInt(fields["userId"], 10, 64)
	return u
}

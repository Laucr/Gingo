package controller

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

func MapConvertToUser(fields map[string]interface{}) *Users {
	u := new(Users)
	u.username = fields["username"].(string)
	u.password = fields["password"].(string)
	u.email = fields["email"].(string)
	u.userId = fields["userId"].(int64)
	return u
}

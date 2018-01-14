package controller

func CreateSession(userName string) int {

	val, _ := Lookup(DbSession, userName)
	if val != nil {
		return SessionExists
	} else {

		Insert()
	}
	return OperationSuccess
}

func CheckSession(userName string) int {

	return OperationSuccess
}
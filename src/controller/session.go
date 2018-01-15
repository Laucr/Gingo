package controller

func CreateSession(userName string) int {

	val, _ := RedisLookup(DbSession, userName)
	if val != nil {
		return SessionExists
	}
	return OperationSuccess
}

func CheckSession(userName string) int {

	return OperationSuccess
}
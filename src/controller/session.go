package controller

const (
	DefaultSessionId = 0
)

func CreateSession(uid int) int {

	//val, _ := RedisLookup(DbSession, userName)
	//if val != nil {
	//	return SessionExists
	//}
	return OperationSuccess
}

func CheckSession(sessionId int) int {

	return OperationSuccess
}
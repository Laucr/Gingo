package controller

import (
	"time"
	"strconv"
)

const (
	DefaultSessionId = ""
)

func CreateSession(userInfo string) (string, int) {
	sessionId := strconv.FormatInt(time.Now().Unix(), 64)
	expireTime := 24 * time.Hour

	insErr, opErr := RedisInsert(DbSession,
		sessionId,
		userInfo,
		expireTime)
	if insErr != InsertSuccess || opErr != OperationSuccess {
		return DefaultSessionId, OperationFailed
	}
	return sessionId, OperationSuccess
}

func GetSession(sessionId string) (*string, int) {
	j, err := RedisSelect(DbSession, sessionId)
	if j == nil || err != OperationSuccess {
		return nil, err
	}
	return j, OperationSuccess
}

func DestorySession(sessionId string) int {
	return OperationSuccess
}

func ExpireSession(sessionId string) int {
	return OperationSuccess
}
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
	ttl, _ := CheckExpire(DbSession, sessionId)
	if ttl < 0 {
		return nil, SessionExpired
	}
	j, err := RedisSelect(DbSession, sessionId)
	if j == nil || err != OperationSuccess {
		return nil, err
	}
	return j, OperationSuccess
}

func ExpireSession(sessionId string) int {
	ttl, _ := CheckExpire(DbSession, sessionId)
	if ttl < 0 {
		return SessionExpired
	}
	res, err := ExpireKey(DbSession, sessionId)
	if res == false || err != OperationSuccess {
		return OperationFailed
	}
	return OperationSuccess
}

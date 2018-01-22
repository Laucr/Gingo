package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	LoginFailed         = -20
	LoginSuccess        = 0
)

func LoginByTel(c *gin.Context) {
	tel := c.PostForm("Tel")
	password := c.PostForm("Password")
	uid, err := GetUid("tel", tel)
	if err != OperationSuccess {
		c.JSON(http.StatusOK, gin.H{
			"status": LoginFailed,
			"error":  err,
		})
		return
	}

	err, sessionId := login(uid, password)
	if err != OperationSuccess {
		c.JSON(http.StatusOK, gin.H{
			"status": LoginFailed,
			"error":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     err,
		"session_id": sessionId,
	})
}

func login(uid int, password string) (int, string) {
	// compare with database
	if CheckPassword(uid, password) == PasswordInvalid {
		return PasswordInvalid, DefaultSessionId
	}

	basicInfo, advInfo, err := GetUserInfo(uid)
	if err < 0  {
		return err, DefaultSessionId
	}

	jUser := ObjConvertToJson(basicInfo) + ObjConvertToJson(advInfo)
	sessionId, _ := CreateSession(jUser)

	return err, sessionId

}

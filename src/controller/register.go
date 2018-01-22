package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"strings"
)

func Register(c *gin.Context) {
	user := new(Users)
	user.UserId = generateUserId()

	user.BasicInfo = new(UserBasicInfo)
	user.BasicInfo.UserId = user.UserId

	user.AdvInfo = new(UserAdvInfo)

	user.Password = c.PostForm("Password")

	user.BasicInfo.UserName = c.PostForm("UserName")
	user.BasicInfo.Email = c.PostForm("Email")
	user.BasicInfo.Tel = c.PostForm("Tel")
	user.BasicInfo.CreateTime = int(time.Now().Unix())

	if InsertUserBasicInfo(user.BasicInfo, user.Password) == InsertSuccess {

		jUser := ObjConvertToJson(user.BasicInfo) + ObjConvertToJson(user.AdvInfo)

		//join 2 json strings
		jUser = strings.Replace(jUser, "}{", ",", -1)
		sessionId, _ := CreateSession(jUser)

		c.JSON(http.StatusOK, gin.H{
			"status":    InsertSuccess,
			"userId":    user.UserId,
			"sessionId": sessionId,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": InsertFailed})
	}
}

func CheckEmailExistence(email string, c *gin.Context) {
	if checkUid, err := GetUid("Email", email); err != OperationSuccess || checkUid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": EmailExists,
			"error":  err,
		})
	}
}

func CheckTelExistence(tel string, c *gin.Context) {
	if checkUid, err := GetUid("Email", tel); err != OperationSuccess || checkUid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": TelExists,
			"error":  err,
		})
	}
}

func generateUserId() int {
	t := time.Now()
	var base = 1514736000
	userId := int(t.Unix()) - base
	return userId
}

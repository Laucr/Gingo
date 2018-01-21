package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	user := new(UserBasicInfo)
	user.UserId = generateUserId()
	user.Password = c.PostForm("Password")
	user.UserName = c.PostForm("UserName")
	user.Email = c.PostForm("Email")
	user.Tel = c.PostForm("Tel")
	user.CreateTime = int(time.Now().Unix())

	if InsertUserBasicInfo(user) == InsertSuccess {
		c.JSON(http.StatusOK, gin.H{
			"status": InsertSuccess,
			"userId": user.UserId})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": InsertFailed})
	}
}

func CheckEmailExistence(email string, c *gin.Context) {
	if checkUid, err := SelectUserBasicInfo("Email", email); err != OperationSuccess || checkUid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": EmailExists,
			"error": err,
		})
	}
}

func CheckTelExistence(tel string, c *gin.Context) {
	if checkUid, err := SelectUserBasicInfo("Email", tel); err != OperationSuccess || checkUid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": TelExists,
			"error": err,
		})
	}
}

func generateUserId() int {
	t := time.Now()
	var base = 1514736000
	userId := int(t.Unix()) - base
	return userId
}

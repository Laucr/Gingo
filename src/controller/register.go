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

	c.JSON(http.StatusOK, gin.H{"status": TelExists})
	c.JSON(http.StatusOK, gin.H{"status": EmailExists})

	if InsertUserBasicInfo(user) == InsertSuccess {
		c.JSON(http.StatusOK, gin.H{
			"status": InsertSuccess,
			"userId": user.UserId})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": InsertFailed})
	}
}

func generateUserId() int {
	t := time.Now()
	var base = 1514736000
	userId := int(t.Unix()) - base
	return userId
}

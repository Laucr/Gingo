package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

const (
	LoginFailed  = -1
	LoginSuccess = 0
)

func PostLogin(c *gin.Context) {
	//user := new(Users)
	//user.UserName = c.PostForm("UserName")
	//user.Password = c.PostForm("Password")
	//userSomeone, err := login(user.UserName, user.Password)
	//if err == LoginSuccess {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status":   LoginSuccess,
	//		"UserName": userSomeone.UserName,
	//		"UserId":   userSomeone.UserId})
	//} else {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": LoginFailed})
	//}

}

func login(uid int, password string) (int, int) {
	// compare with database
	if CheckPassword(uid, password) == PasswordInvalid {
		return LoginFailed, DefaultSessionId
	}
	return LoginSuccess, CreateSession(uid)

}

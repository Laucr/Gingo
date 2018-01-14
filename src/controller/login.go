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
	user := new(Users)
	user.userName = c.PostForm("userName")
	user.password = c.PostForm("password")
	userSomeone, err := login(user.userName, user.password)
	if err == LoginSuccess {
		c.JSON(http.StatusOK, gin.H{
			"status":   LoginSuccess,
			"userName": userSomeone.userName,
			"userId":   userSomeone.userId})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": LoginFailed})
	}

}

func login(username string, password string) (*Users, int) {
	// compare with database
	val, err := Lookup(DbUserInfo, username)
	if val == nil {
		fmt.Println("Error:", err)
		return nil, LoginFailed
	}
	userSomeone := MapConvertToUser(*val)
	if userSomeone.userName == username && userSomeone.password == password {
		return userSomeone, LoginSuccess
	} else {
		return nil, LoginFailed
	}
}

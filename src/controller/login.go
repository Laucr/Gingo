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
	user.UserName = c.PostForm("UserName")
	user.Password = c.PostForm("Password")
	userSomeone, err := login(user.UserName, user.Password)
	if err == LoginSuccess {
		c.JSON(http.StatusOK, gin.H{
			"status":   LoginSuccess,
			"UserName": userSomeone.UserName,
			"UserId":   userSomeone.UserId})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": LoginFailed})
	}

}

func login(username string, password string) (*Users, int) {
	// compare with database
	val, err := RedisLookup(DbUsers, username)
	if val == nil {
		fmt.Println("Error:", err)
		return nil, LoginFailed
	}
	userSomeone := MapConvertToUser(*val)
	if userSomeone.UserName == username && userSomeone.Password == password {
		return userSomeone, LoginSuccess
	} else {
		return nil, LoginFailed
	}
}

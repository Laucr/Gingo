package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
)

func Register(c *gin.Context) {
	user := new(Users)
	user.userName = c.PostForm("userName")
	user.password = c.PostForm("password")
	user.userId = generateUserId()

	if generateUser(user) == InsertSuccess {
		//c.String(http.StatusOK, "Hello, %s", user.userName)
		c.JSON(http.StatusOK, gin.H{
			"status":   InsertSuccess,
			"username": user.userName})
	} else {
		//c.String(http.StatusOK, "Sorry, failed to register")
		c.JSON(http.StatusOK, gin.H{
			"status": InsertFailed})
	}

}

func generateUserId() int64 {
	t := time.Now()
	var base int64 = 1514736000
	userId := t.Unix() - base
	return userId
}

func generateUser(user *Users) int {
	// user -> database
	u := UsersConvertToMap(user)

	// insert_result function_result
	insRes, funcRes := Insert(DbUserInfo, user.userName, u)
	if funcRes != OperationSuccess {
		fmt.Println("Error:", insRes)
	}

	return insRes
}

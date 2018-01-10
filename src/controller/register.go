package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
)

func Register(c *gin.Context) {
	user := new(Users)
	user.username = c.PostForm("username")
	user.email = c.PostForm("email")
	user.password = c.PostForm("password")
	user.userId = generateUserId()
	if generateUser(user) == InsertSuccess {
		c.String(http.StatusOK, "Hello, %s", user.username)
	} else {
		c.String(http.StatusOK, "Sorry, failed to register")
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
	insr, funcr := Insert(UserInfo, user.username, u)
	if funcr != OperationSuccess {
		fmt.Println("Error:", insr)
	}

	return insr
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	user := new(Users)
	user.username = c.PostForm("username")
	user.email = c.PostForm("email")
	user.password = c.PostForm("password")
	user.userId = generateUserId()
	genetateUser(user)
	c.String(http.StatusOK, "Hello, %s", user.username)

}

func generateUserId() int64 {
	return 0
}

func genetateUser(user *Users) int{
	// user -> database
	return 0
}
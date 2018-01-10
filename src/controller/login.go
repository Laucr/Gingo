package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostLogin(c *gin.Context) {
	user := new(Users)
	user.username = c.PostForm("username")
	user.password = c.PostForm("password")
	if login(user.username, user.password) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":   "Login",
			"username": user.username,
			"email":    user.email,
			"userId":   user.userId})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "error"})
	}

}

func login(username string, upwd string) int {
	// compare with database

	return 0
}

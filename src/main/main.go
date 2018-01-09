package main

import (
	"github.com/gin-gonic/gin"
	"controller"
)

func main() {
	router := gin.Default()
	router.POST("/register/submit", controller.Register)
	router.POST("/login/submit", controller.PostLogin)
	router.Run(":51234")
}


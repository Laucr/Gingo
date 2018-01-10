package main

import (
	"github.com/gin-gonic/gin"
	"controller"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("src/main/*.html")
	router.POST("/register/submit", controller.Register)
	router.POST("/login/submit", controller.PostLogin)
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index",
			"body": "hello gingo"})
	})
	router.Run(":51234")
}


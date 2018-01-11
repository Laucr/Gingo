package main

import (
	"github.com/gin-gonic/gin"
	"controller"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()
	workDir, _ := filepath.Abs(".")
	router.LoadHTMLGlob(filepath.Join(workDir, "../templates/*.html"))
	router.POST("/register/submit", controller.Register)
	router.POST("/login/submit", controller.PostLogin)
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index",
			"body": "hello gingo"})
	})
	router.Run(":51234")
}


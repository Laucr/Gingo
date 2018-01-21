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

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index",
			"body":  "hello gingo"})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	router.POST("/register/submit", controller.Register)
	router.POST("/login/submit", controller.PostLogin)
	router.Run(":51234")
}

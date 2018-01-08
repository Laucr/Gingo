package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gingo!")
	})
	router.Run(":51234")
}



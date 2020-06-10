package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	gin.SetMode("debug")

	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}

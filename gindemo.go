package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.GetStringMap("gin ping")
		c.JSON(200, gin.H{
			"message": "ping ok",
		})

	})

	app.Run()
}

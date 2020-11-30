package main

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
)

//import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "hello",
		})
	})

	r.GET("/testdb", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "testdb",
		})
	})

	r.Group("/hello", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

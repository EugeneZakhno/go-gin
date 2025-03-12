package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Create a routing group
	v1 := router.Group("/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.String(200, "User list in v1")
		})
		v1.GET("/posts", func(c *gin.Context) {
			c.String(200, "Post list in v1")
		})
	}
	router.Run(":8080")
}

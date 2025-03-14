package main

import (
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simulate authentication
		if c.GetHeader("Authorization") != "secret" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func main() {
	router := gin.Default()
	// Add Logger and Recovery middleware globally
	// Create a routing group, and all routes in this group will apply the authMiddleware middleware
	authenticated := router.Group("/")
	authenticated.Use(authMiddleware())
	{
		authenticated.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})

		authenticated.GET("/private", func(c *gin.Context) {
			c.String(200, "Private data")
		})
	}

	// Not in the routing group, so the authMiddleware middleware is not applied
	router.GET("/welcome", func(c *gin.Context) {
		c.String(200, "Welcome!")
	})

	router.Run(":8080")
}

package app

import "github.com/gin-gonic/gin"

// Routes sets up the routes for the application.
func Routes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}

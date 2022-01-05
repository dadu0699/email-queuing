package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status is the status of the application.
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

package handlers

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(r *gin.Engine) {
	// Simple ping route
	r.GET("/ping", PingHandler)
}

// PingHandler handles the /ping route
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

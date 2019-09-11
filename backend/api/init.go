package api

import (
	"github.com/gin-gonic/gin"
)

// InitAPI xxx
func InitAPI() *gin.Engine {
	router := gin.Default()

	router.HEAD("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", signUp)
		auth.POST("/login", login)
	}

	school := router.Group("/school")
	{
		school.GET("/class", classList)
		school.POST("/class", addClass)
	}
	return router
}

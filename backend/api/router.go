package api

import (
	"Microservice/backend/middleware"

	"github.com/gin-gonic/gin"
)

/**
注册: curl -XPOST http://localhost:8080/auth/sign_up -d 'name=ysy&account=as@sina.com&password=123'
登录: curl -XPOST http://localhost:8080/auth/login -d'account=as@sina.com&password=123'

获取课程(需登录): curl -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50IjoiYXNAc2luYS5jb20iLCJleHAiOjE1NjgxMDY4NjMsInVzZXJOYW1lIjoieXN5In0.sNgq4G2GdKvoDaijIosCWNn5RriDPn2VL-qcxjaR9E0" -XGET http://localhost:8080/school/class

*/

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

	// 需要验证登录
	school := router.Group("/school")
	{
		school.GET("/class", classList)
		school.POST("/class", addClass)
	}
	school.Use(middleware.AuthMiddleware())
	return router
}

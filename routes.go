// 路由处理
package main

import (
	"assemblyline/project1023/controller"
	"assemblyline/project1023/middleware"

	"github.com/gin-gonic/gin"
)

// 定义路由
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}

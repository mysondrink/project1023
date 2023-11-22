// 路由处理
package main

import (
	"assemblyline/project1023/controller"
	"assemblyline/project1023/middleware"

	"github.com/gin-gonic/gin"
)

// 定义路由
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.GET("/api/trail/info", controller.GetInfo)
	r.GET("/api/module/info", controller.GetInfo)
	r.GET("/api/car/info", controller.GetCarInfo)
	r.GET("/api/car/column", controller.GetCarInfo)
	r.GET("/api/layout/info", controller.GetToyInfo)
	return r
}

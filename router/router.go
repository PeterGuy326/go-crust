package router

import (
	"crust/controller"
	"crust/middleware"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {

	// router group
	group := r.Group("/api/v1.0.0")

	// 全局加载日志中间件
	group.Use(middleware.RecordPostLog())

	// 用户模块
	group.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": controller.Login(c),
		})
	})
	group.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": controller.Register(c),
		})
	})

}

package main

import (
	"crust/config"
	"crust/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化配置文件
	config.ViperInit()

	// 初始化mysql连接
	config.InitMysqlDB()

	// 路由定义
	router.Route(r)

	_ = r.Run()
}

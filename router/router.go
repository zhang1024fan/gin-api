// utils/router.go

package router

import (
	"gin-api/api/system/controller"
	"gin-api/common/config"
	"gin-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// 导入模块路由注册器
	"gin-api/router/system"
	//"gin-api/utils/k8s"
)

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	// 中间件
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	// 让 /upload/ 开头的请求映射到 ./upload/ 目录下
	router.Static("/upload", config.Config.ImageSettings.UploadDir)
	router.Use(middleware.Logger())

	// 路由注册
	register(router)

	return router
}

// 路由注册中心
func register(router *gin.Engine) {
	// 公共接口：Swagger、静态资源等
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.Static("/upload", config.Config.ImageSettings.UploadDir)

	// 统一 API 分组
	apiGroup := router.Group("/api")
	{
		// 不需要 JWT 的接口
		apiGroup.GET("/captcha", controller.Captcha)
		apiGroup.POST("/login", controller.Login)
		// 需要 JWT鉴权 的接口
		jwtGroup := apiGroup.Group("")
		jwtGroup.Use(middleware.AuthMiddleware(), middleware.LogMiddleware())
		{
			system.RegisterSystemRoutes(jwtGroup)
		}
	}
}

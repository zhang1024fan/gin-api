// utils/router.go

package router

import (
	"gin-api/api/system/controller"
	"gin-api/common/config"
	"gin-api/middleware"
	"gin-api/pkg/log"

	"gin-api/router/k8s"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"reflect"
	"strings"

	"gin-api/router/cmdb"         // cmdb模块路由
	"gin-api/router/configCenter" // 配置中心模块路由
	"gin-api/router/system"       // 系统模块路由
)

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 告诉 validator 使用 json 标签作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			tag := fld.Tag.Get("json")
			if tag == "-" {
				return ""
			}
			// 去掉 json 标签里的 ",omitempty"
			return strings.Split(tag, ",")[0]
		})
	}

	// 中间件
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	// 让 /upload/ 开头的请求映射到 ./upload/ 目录下
	router.Static("/upload", config.Config.ImageSettings.UploadDir)
	router.Use(log.CustomGinLogger())

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
	apiGroup := router.Group("/api/v1")
	{
		// 不需要 JWT 的接口
		apiGroup.GET("/captcha", controller.Captcha)
		apiGroup.POST("/login", controller.Login)
		// 需要 JWT鉴权 的接口
		jwtGroup := apiGroup.Group("")
		jwtGroup.Use(middleware.AuthMiddleware())
		{
			system.RegisterSystemRoutes(jwtGroup)
			cmdb.RegisterCmdbRoutes(jwtGroup)
			configCenter.RegisterConfigCenterRoutes(jwtGroup) // 新增这一行
			k8s.RegisterK8sRoutes(jwtGroup)
		}
	}
}

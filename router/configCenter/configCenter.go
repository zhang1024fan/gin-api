// configCenter路由注册系统

package configCenter

import (
	"gin-api/api/config_center/controller"

	"github.com/gin-gonic/gin"
)

// RegisterConfigCenterRoutes 注册配置中心相关路由
func RegisterConfigCenterRoutes(router *gin.RouterGroup) {
	ecsAuthCtrl := controller.NewEcsAuthController()

	// ECS认证凭据管理
	router.GET("/config/ecsauthlist", ecsAuthCtrl.GetEcsAuthList)     // 获取所有凭据
	router.GET("/config/ecsauthinfo", ecsAuthCtrl.GetEcsAuthByName)   // 根据名称获取凭据
	router.POST("/config/ecsauthadd", ecsAuthCtrl.CreateEcsAuth)      // 创建凭据
	router.PUT("/config/ecsauthupdate", ecsAuthCtrl.UpdateEcsAuth)    // 更新凭据
	router.DELETE("/config/ecsauthdelete", ecsAuthCtrl.DeleteEcsAuth) // 删除凭据
}

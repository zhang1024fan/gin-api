// cmdb路由注册系统

package cmdb

import (
	"gin-api/api/cmdb/controller"

	"github.com/gin-gonic/gin"
)

// RegisterCmdbRoutes 注册系统相关路由
func RegisterCmdbRoutes(router *gin.RouterGroup) {
	// 资产分组
	router.POST("/cmdb/groupadd", controller.CreateCmdbGroup)      // 添加资产分组
	router.GET("/cmdb/grouplist", controller.GetAllCmdbGroups)     // 获取所有资产分组
	router.PUT("/cmdb/groupupdate", controller.UpdateCmdbGroup)    // 更新资产分组
	router.DELETE("/cmdb/groupdelete", controller.DeleteCmdbGroup) // 删除资产分组
	router.GET("/cmdb/groupbyname", controller.GetCmdbGroupByName) // 根据名称查询分组
}

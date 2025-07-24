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

	// 主机管理
	router.POST("/cmdb/hostcreate", controller.NewCmdbHostController().CreateCmdbHost)       // 创建主机
	router.PUT("/cmdb/hostupdate", controller.NewCmdbHostController().UpdateCmdbHost)    // 更新主机
	router.DELETE("/cmdb/hostdelete", controller.NewCmdbHostController().DeleteCmdbHost) // 删除主机
	router.GET("/cmdb/hostlist", controller.NewCmdbHostController().GetCmdbHostListWithPage) // 获取主机列表(分页)
	router.GET("/cmdb/hostinfo", controller.NewCmdbHostController().GetCmdbHostById)     // 根据ID获取主机
	router.GET("/cmdb/hostgroup", controller.NewCmdbHostController().GetCmdbHostsByGroupId) // 根据分组ID获取主机列表
	router.GET("/cmdb/hostbyname", controller.NewCmdbHostController().GetCmdbHostsByHostNameLike) // 根据主机名称模糊查询
	router.GET("/cmdb/hostbyip", controller.NewCmdbHostController().GetCmdbHostsByIP)     // 根据IP查询主机
	router.GET("/cmdb/hostbystatus", controller.NewCmdbHostController().GetCmdbHostsByStatus) // 根据状态查询主机
}

package controller

import (
	"gin-api/api/cmdb/model"
	"gin-api/api/cmdb/service"
	"gin-api/common/constant"
	"gin-api/common/result"

	"github.com/gin-gonic/gin"
)

var sysDept model.CmdbGroup

// @Summary 新增资产分组接口
// @Produce json
// @Description 新增资产分组接口
// @Param data body model.CmdbGroup true "data"
// @Success 200 {object} result.Result
// @router /api/cmdb/groupadd [post]
// @Security ApiKeyAuth
func CreateCmdbGroup(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.GetCmdbGroupService().CreateCmdbGroup(c, sysDept)
}

// @Summary 查询所有资产分组（树形结构）
// @Produce json
// @Description 查询所有资产分组，并以树形结构返回
// @Success 200 {object} result.Result
// @router /api/cmdb/grouplist [get]
// @Security ApiKeyAuth
func GetAllCmdbGroups(c *gin.Context) {
	service.GetCmdbGroupService().GetAllCmdbGroups(c)
}

// @Summary 更新资产分组接口
// @Produce json
// @Description 更新资产分组接口
// @Param data body model.CmdbGroup true "data"
// @Success 200 {object} result.Result
// @router /api/cmdb/groupupdate [put]
// @Security ApiKeyAuth
func UpdateCmdbGroup(c *gin.Context) {
	var group model.CmdbGroup
	_ = c.BindJSON(&group)
	service.GetCmdbGroupService().UpdateCmdbGroup(c, group)
}

// @Summary 删除资产分组接口
// @Produce json
// @Description 删除资产分组接口
// @Param data body model.CmdbGroupIdDto true "分组ID"
// @Success 200 {object} result.Result
// @router /api/cmdb/groupdelete [delete]
// @Security ApiKeyAuth
func DeleteCmdbGroup(c *gin.Context) {
	var dto model.CmdbGroupIdDto
	if err := c.BindJSON(&dto); err != nil {
		result.Failed(c, constant.GROUP_EXIST, "参数错误")
		return
	}
	service.GetCmdbGroupService().DeleteCmdbGroup(c, dto.Id)
}

package controller

import (
	"gin-api/api/cmdb/model"
	"gin-api/api/cmdb/service"
	"gin-api/common/constant"
	"gin-api/common/result"
	"gin-api/common/util"

	"github.com/gin-gonic/gin"
)

type CmdbHostController struct {
	service service.CmdbHostServiceInterface
}

func NewCmdbHostController() *CmdbHostController {
	return &CmdbHostController{
		service: service.GetCmdbHostService(),
	}
}

// 获取主机列表
// @Summary 获取主机列表
// @Description 获取主机列表
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostlist [get]
// @Security ApiKeyAuth
func (c *CmdbHostController) GetCmdbHostList(ctx *gin.Context) {
	c.service.GetCmdbHostList(ctx)
}

// 创建主机
// @Summary 创建主机
// @Description 创建主机
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param data body model.CreateCmdbHostDto true "主机信息"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostcreate [post]
// @Security ApiKeyAuth
func (c *CmdbHostController) CreateCmdbHost(ctx *gin.Context) {
	var dto model.CreateCmdbHostDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		result.Failed(ctx, constant.INVALID_PARAMS, "参数错误")
		return
	}
	c.service.CreateCmdbHost(ctx, &dto)
}

// 更新主机
// @Summary 更新主机
// @Description 更新主机
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param data body model.UpdateCmdbHostDto true "主机信息(包含ID)"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostupdate [put]
// @Security ApiKeyAuth
func (c *CmdbHostController) UpdateCmdbHost(ctx *gin.Context) {
	var dto model.UpdateCmdbHostDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		result.Failed(ctx, constant.INVALID_PARAMS, "参数错误")
		return
	}
	c.service.UpdateCmdbHost(ctx, dto.ID, &dto)
}

// 删除主机
// @Summary 删除主机
// @Description 删除主机
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param data body model.CmdbHostIdDto true "主机ID"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostdelete [delete]
// @Security ApiKeyAuth
func (c *CmdbHostController) DeleteCmdbHost(ctx *gin.Context) {
	var dto model.CmdbHostIdDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		result.Failed(ctx, constant.INVALID_PARAMS, "参数错误")
		return
	}
	c.service.DeleteCmdbHost(ctx, dto.ID)
}

// 根据ID获取主机
// @Summary 根据ID获取主机
// @Description 根据ID获取主机
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param id query int true "主机ID"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostinfo [get]
// @Security ApiKeyAuth
func (c *CmdbHostController) GetCmdbHostById(ctx *gin.Context) {
	id := util.StringToUint(ctx.Query("id"))
	c.service.GetCmdbHostById(ctx, id)
}

// 根据分组ID获取主机列表
// @Summary 根据分组ID获取主机列表
// @Description 根据分组ID获取主机列表
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param groupId query int true "分组ID"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostgroup [get]
// @Security ApiKeyAuth
func (c *CmdbHostController) GetCmdbHostsByGroupId(ctx *gin.Context) {
	groupId := util.StringToUint(ctx.Query("groupId"))
	c.service.GetCmdbHostsByGroupId(ctx, groupId)
}

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

// 分页参数
type PageParams struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=1,max=100"`
}

// 获取主机列表(分页)
// @Summary 获取主机列表(分页)
// @Description 获取主机列表(分页)
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param page query int true "页码"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostlist [get]
// @Security ApiKeyAuth
func (c *CmdbHostController) GetCmdbHostListWithPage(ctx *gin.Context) {
	var params PageParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		result.Failed(ctx, constant.INVALID_PARAMS, "分页参数错误")
		return
	}
	c.service.GetCmdbHostListWithPage(ctx, params.Page, params.PageSize)
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

// 根据主机名称模糊查询
// @Summary 根据主机名称模糊查询
// @Description 根据主机名称模糊查询
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param name query string true "主机名称(模糊匹配)"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostbyname [get]
// @Security ApiKeyAuth
func (c *CmdbHostController) GetCmdbHostsByHostNameLike(ctx *gin.Context) {
	name := ctx.Query("name")
	c.service.GetCmdbHostsByHostNameLike(ctx, name)
}

// 根据IP查询主机
// @Summary 根据IP查询主机
// @Description 根据IP查询主机(匹配内网IP、公网IP或SSH IP)
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param ip query string true "IP地址"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostbyip [get]
// @Security ApiKeyAuth
func (c *CmdbHostController) GetCmdbHostsByIP(ctx *gin.Context) {
	ip := ctx.Query("ip")
	c.service.GetCmdbHostsByIP(ctx, ip)
}

// 根据状态查询主机
// @Summary 根据状态查询主机
// @Description 根据状态查询主机(1->认证成功,2->未认证,3->认证失败)
// @Tags CMDB资产管理
// @Accept json
// @Produce json
// @Param status query int true "状态(1/2/3)"
// @Success 200 {object} result.Result
// @Router /api/v1/cmdb/hostbystatus [get]
// @Security ApiKeyAuth
func (c *CmdbHostController) GetCmdbHostsByStatus(ctx *gin.Context) {
	status := int(util.StringToUint(ctx.Query("status")))
	c.service.GetCmdbHostsByStatus(ctx, status)
}

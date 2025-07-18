package controller

import (
	"gin-api/api/config_center/model"
	"gin-api/api/config_center/service"
	"gin-api/common/result"

	"github.com/gin-gonic/gin"
)

type EcsAuthController struct {
	service service.EcsAuthServiceInterface
}

func NewEcsAuthController() *EcsAuthController {
	return &EcsAuthController{
		service: service.GetEcsAuthService(),
	}
}

// GetEcsAuthList 获取所有凭据
// @Summary 获取所有凭据
// @Tags Config配置中心
// @Success 200 {object} result.Result{data=[]model.EcsAuthVo}
// @Router /api/v1/config/ecsauthlist [get]
// @Security ApiKeyAuth
func (c *EcsAuthController) GetEcsAuthList(ctx *gin.Context) {
	c.service.GetEcsAuthList(ctx)
}

// CreateEcsAuth 创建凭据
// @Summary 创建凭据
// @Tags Config配置中心
// @Param data body model.CreateEcsPasswordAuthDto true "凭据信息"
// @Success 200 {object} result.Result
// @Router /api/v1/config/ecsauthadd [post]
// @Security ApiKeyAuth
func (c *EcsAuthController) CreateEcsAuth(ctx *gin.Context) {
	var dto model.CreateEcsPasswordAuthDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		result.Failed(ctx, int(result.ApiCode.FAILED), err.Error())
		return
	}
	c.service.CreateEcsAuth(ctx, &dto)
}

// GetEcsAuthByName 根据名称获取凭据
// @Summary 根据名称获取凭据
// @Tags Config配置中心
// @Param name query string true "凭据名称"
// @Success 200 {object} result.Result{data=model.EcsAuthVo}
// @Router /api/v1/config/ecsauthinfo [get]
// @Security ApiKeyAuth
func (c *EcsAuthController) GetEcsAuthByName(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		result.Failed(ctx, int(result.ApiCode.FAILED), "name参数不能为空")
		return
	}
	c.service.GetEcsAuthByName(ctx, name)
}

// UpdateEcsAuth 更新凭据
// @Summary 更新凭据
// @Tags Config配置中心
// @Param data body model.UpdateEcsAuthDto true "凭据信息"
// @Success 200 {object} result.Result
// @Router /api/v1/config/ecsauthupdate [put]
// @Security ApiKeyAuth
func (c *EcsAuthController) UpdateEcsAuth(ctx *gin.Context) {
	var dto model.UpdateEcsAuthDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		result.Failed(ctx, int(result.ApiCode.FAILED), err.Error())
		return
	}
	c.service.UpdateEcsAuth(ctx, dto.Id, &dto.CreateEcsPasswordAuthDto)
}

// DeleteEcsAuth 删除凭据
// @Summary 删除凭据
// @Tags Config配置中心
// @Param data body model.EcsAuthIdDto true "凭据ID"
// @Success 200 {object} result.Result
// @Router /api/v1/config/ecsauthdelete [delete]
// @Security ApiKeyAuth
func (c *EcsAuthController) DeleteEcsAuth(ctx *gin.Context) {
	var idDto model.EcsAuthIdDto
	if err := ctx.ShouldBindJSON(&idDto); err != nil {
		result.Failed(ctx, int(result.ApiCode.FAILED), err.Error())
		return
	}
	c.service.DeleteEcsAuth(ctx, idDto.Id)
}

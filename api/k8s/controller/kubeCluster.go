package controller

import (
	"errors"
	"gin-api/api/k8s/model"
	"gin-api/api/k8s/service"
	"gin-api/common/result"
	"gin-api/common/valid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Tags k8s集群列表
// @Summary k8s集群列表接口
// @Produce json
// @Description k8s集群列表接口
// @Success 200 {object} result.Result
// @router /api/v1/clusters [get]
// @Security ApiKeyAuth
func QueryCluster(c *gin.Context) {
	service.KubeClusterService().QueryClusterList(c)
}

// @Tags k8s集群列表
// @Summary 创建Kubernetes集群
// @Description 根据提供的配置信息创建一个新的Kubernetes集群
// @Accept json
// @Produce json
// @Param cluster body model.AddKubeClusterDto true "集群配置信息"
// @Success 200 {object} result.Result "操作成功返回结果"
// @Router /api/v1/clusters/add [post]
// @Security ApiKeyAuth
func CreateCluster(c *gin.Context) {
	var dto model.AddKubeClusterDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 返回校验的错误
			result.Failed(c, int(result.ApiCode.ValidationParameterError), result.ApiCode.GetMessage(result.ApiCode.ValidationParameterError)+":"+valid.ErrorToText(ve))
		}
		return
	}
	service.KubeClusterService().CreateCluster(c, dto)
}

// @Tags k8s集群列表
// @Summary 更新Kubernetes集群信息
// @Description 根据提供的更新数据修改已有Kubernetes集群的配置
// @Accept json
// @Produce json
// @Param cluster body model.UpdateKubeClusterDto true "更新后的集群配置信息"
// @Success 200 {object} result.Result "操作成功返回结果"
// @Router /api/v1/clusters/update [post]
// @Security ApiKeyAuth
func UpdateCluster(c *gin.Context) {
	var dto model.UpdateKubeClusterDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 返回校验的错误
			result.Failed(c, int(result.ApiCode.ValidationParameterError), result.ApiCode.GetMessage(result.ApiCode.ValidationParameterError)+":"+valid.ErrorToText(ve))
		}
		return
	}
	service.KubeClusterService().UpdateCluster(c, dto)
}

package service

import (
	"gin-api/api/k8s/dao"
	"gin-api/api/k8s/model"
	"gin-api/common/result"
	"gin-api/common/util"
	"github.com/gin-gonic/gin"
	"time"
)

type IKubeClusterService interface {
	QueryClusterList(c *gin.Context)                              // 查询k8s集群列表
	CreateCluster(c *gin.Context, dto model.AddKubeClusterDto)    // 查询k8s集群列表
	UpdateCluster(c *gin.Context, dto model.UpdateKubeClusterDto) // 查询k8s集群列表
}
type KubeClusterServiceImpl struct{}

// 集群下拉列表
func (s KubeClusterServiceImpl) QueryClusterList(c *gin.Context) {
	result.Success(c, dao.QueryKubeClusterVoList())
}

// 添加集群配置信息
func (s KubeClusterServiceImpl) CreateCluster(c *gin.Context, dto model.AddKubeClusterDto) {
	kubeClusterByName := dao.GetKubeClusterByName(dto.Name)
	if kubeClusterByName.ID > 0 {
		result.Failed(c, int(result.ApiCode.KUBEClUSTEREXIST), result.ApiCode.GetMessage(result.ApiCode.KUBEClUSTEREXIST))
		return
	}

	model := model.KubeCluster{
		Name:       dto.Name,
		KubeConfig: dto.KubeConfig,
		IsDefault:  dto.IsDefault,
		CreateTime: util.HTime{Time: time.Now()},
		UpdateTime: util.HTime{Time: time.Now()},
	}

	if !dao.AddKubeCluster(&model) {
		result.Failed(c, int(result.ApiCode.FAILED), result.ApiCode.GetMessage(result.ApiCode.FAILED))
		return
	}

	result.Success(c, true)
}

// 修改集群配置信息
func (s KubeClusterServiceImpl) UpdateCluster(c *gin.Context, dto model.UpdateKubeClusterDto) {
	kubeClusterByName := dao.GetKubeClusterByID(dto.ID)
	if kubeClusterByName.ID == 0 {
		result.Failed(c, int(result.ApiCode.KUBEClUSTERNOTEXIST), result.ApiCode.GetMessage(result.ApiCode.KUBEClUSTERNOTEXIST))
		return
	}

	model := model.KubeCluster{
		ID:         dto.ID,
		KubeConfig: dto.KubeConfig,
		IsDefault:  dto.IsDefault,
		UpdateTime: util.HTime{Time: time.Now()},
	}

	if !dao.UpdateKubeCluster(&model) {
		result.Failed(c, int(result.ApiCode.FAILED), result.ApiCode.GetMessage(result.ApiCode.FAILED))
		return
	}

	result.Success(c, true)
}

var kubeClusterService = KubeClusterServiceImpl{}

func KubeClusterService() IKubeClusterService {
	return &kubeClusterService
}

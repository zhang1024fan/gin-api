package service

import (
	"encoding/json"
	"fmt"
	"gin-api/api/k8s/dao"
	"gin-api/api/k8s/model"
	"gin-api/common/result"
	"gin-api/common/util"
	"gin-api/pkg/log"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"time"
)

type IKubeClusterService interface {
	QueryClusterList(c *gin.Context)                                              // 查询k8s集群列表
	CreateCluster(c *gin.Context, dto model.AddKubeClusterDto)                    // 查询k8s集群列表
	UpdateCluster(c *gin.Context, dto model.UpdateKubeClusterDto)                 // 查询k8s集群列表
	GetClusterByName(c *gin.Context, name string) (kubeCluster model.KubeCluster) // 通过名称查询集群信息
	SaveUnstructuredListToRedis(key string, uList *unstructured.UnstructuredList) bool
	LoadUnstructuredListFromRedis(c *gin.Context, key string) *unstructured.UnstructuredList
}
type KubeClusterServiceImpl struct{}

func (s KubeClusterServiceImpl) SaveUnstructuredListToRedis(key string, uList *unstructured.UnstructuredList) bool {

	data, err := json.Marshal(uList)
	if err != nil {
		log.Log().Error("marshal error: ", err)
		return false
	}
	return util.RedisStore{}.SetKubeClusterResource(key, data)

}

func (s KubeClusterServiceImpl) LoadUnstructuredListFromRedis(c *gin.Context, key string) *unstructured.UnstructuredList {
	data := util.RedisStore{}.GetKubeClusterResource(key)
	if data == "" {
		log.Log().Error("data is empty or error")
		return nil
	}
	var uList unstructured.UnstructuredList
	err := json.Unmarshal([]byte(data), &uList)
	if err != nil {
		log.Log().Error("unmarshal error: ", err)
		return nil
	}
	return &uList
}

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

	//删除redis缓存
	util.RedisStore{}.DelKubeCluster(kubeClusterByName.Name)

	result.Success(c, true)
}

func (s KubeClusterServiceImpl) GetClusterByName(c *gin.Context, name string) (kubeCluster model.KubeCluster) {

	//从redis中获取
	kubeClusterString := util.RedisStore{}.GetKubeCluster(name)

	if kubeClusterString == "" {
		//从redis中未获取到
		//尝试从数据库中获取
		kubeCluster = dao.GetKubeClusterByName(name)
		//如果获取到，写入到redis中
		if kubeCluster.ID > 0 {
			//将结构体序列化为json
			temp, err := json.Marshal(kubeCluster)
			if err != nil {
				log.Log().Error("redis序列化失败", err)
				return
			}
			_ = util.RedisStore{}.SetKubeCluster(name, string(temp))
		}

		return
	}

	//从redis中获取到
	//1.解析string为json
	if err := json.Unmarshal([]byte(kubeClusterString), &kubeCluster); err != nil {
		fmt.Println("JSON 解析失败:", err)
		return
	}
	return
}

var kubeClusterService = KubeClusterServiceImpl{}

func KubeClusterService() IKubeClusterService {
	return &kubeClusterService
}

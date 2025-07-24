package service

import (
	"gin-api/api/k8s/model"
	"gin-api/api/k8s/resource"
	"gin-api/common/result"
	"gin-api/pkg/log"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"time"
)

var refreshGroup singleflight.Group

type IKubeResourceService interface {
	List(c *gin.Context)
	IsClusterScoped() bool
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
type KubeResourceServiceImpl struct {
	client          schema.GroupVersionResource
	isClusterScoped bool
}

func NewKubeResourceService(client schema.GroupVersionResource, IsClusterScoped bool) IKubeResourceService {
	return &KubeResourceServiceImpl{
		client:          client,
		isClusterScoped: IsClusterScoped,
	}
}

// @Tags K8s集群
// @Summary 获取k8s集群某个资源对象
// @Description 获取k8s集群中某个资源的对象
// @Accept json
// @Produce json
// @Param x-cluster-name header string true "集群名字信息"
// @Success  200  {object}  map[string]interface{}
// @Router /api/v1/k8s/resource/deployments/{namespace}/{name} [get]
// @Security ApiKeyAuth
func (s KubeResourceServiceImpl) Get(c *gin.Context) {
	value, exists := c.Get("cluster")
	if !exists {
		result.Failed(c, int(result.ApiCode.KUBEClUSTERNAMENOTEXIST), result.ApiCode.GetMessage(result.ApiCode.KUBEClUSTERNAMENOTEXIST))
		return
	}
	kubeCluster, ok := value.(*model.KubeCluster)
	if !ok || kubeCluster.Name == "" {
		result.Failed(c, int(result.ApiCode.KUBEClUSTERNAMENOTEXIST), result.ApiCode.GetMessage(result.ApiCode.KUBEClUSTERNAMENOTEXIST))
		return
	}
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeCluster.KubeConfig))
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "kubeconfig error")
	}
	client := resource.NewDynamicResourceClient(s.client, s.isClusterScoped)

	obj := client.Get(config, c.Param("namespace"), c.Param("name"))

	c.JSON(200, obj)
}

// @Tags K8s集群
// @Summary 获取k8s集群某个资源列表
// @Description 获取k8s集群中某个资源的列表
// @Accept json
// @Produce json
// @Param x-cluster-name header string true "集群名字信息"
// @Success  200  {object}  map[string]interface{}
// @Router /api/v1/k8s/resource/deployments/{namespace} [get]
// @Security ApiKeyAuth
func (s KubeResourceServiceImpl) List(c *gin.Context) {
	var objList *unstructured.UnstructuredList
	value, exists := c.Get("cluster")
	if !exists {
		result.Failed(c, int(result.ApiCode.KUBEClUSTERNAMENOTEXIST), result.ApiCode.GetMessage(result.ApiCode.KUBEClUSTERNAMENOTEXIST))
		return
	}
	kubeCluster, ok := value.(*model.KubeCluster)
	if !ok || kubeCluster == nil || kubeCluster.Name == "" {
		result.Failed(c, int(result.ApiCode.KUBEClUSTERNAMENOTEXIST), result.ApiCode.GetMessage(result.ApiCode.KUBEClUSTERNAMENOTEXIST))
		return
	}

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeCluster.KubeConfig))
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "kubeconfig error")
	}
	client := resource.NewDynamicResourceClient(s.client, s.isClusterScoped)

	//尝试从redis缓存中去获取
	cacheKey := kubeCluster.Name + ":" + c.Request.URL.String()
	log.Log().Info("kubeCluster CacheKey: ", cacheKey)
	objList = KubeClusterService().LoadUnstructuredListFromRedis(c, kubeCluster.Name+":"+c.Request.URL.String())
	if objList != nil && objList.IsList() {
		// 在redis中查询到缓存
		log.Log().Info("CacheKey: ", cacheKey, " get cache success")
		c.JSON(http.StatusOK, objList)
		//异步 5s后 自动调用下接口 刷新下缓存
		go func() {
			_, _, shared := refreshGroup.Do(cacheKey, func() (interface{}, error) {
				log.Log().Info("CacheKey: ", cacheKey, " waiting refresh cache")
				time.Sleep(5 * time.Second)
				log.Log().Info("get new data from k8s ", cacheKey)
				objList = client.List(config, c.Param("namespace"))
				//在redis中缓存相关查询
				if !KubeClusterService().SaveUnstructuredListToRedis(kubeCluster.Name+":"+c.Request.URL.String(), objList) {
					log.Log().Warn(kubeCluster.Name + ":" + c.Request.URL.String() + " set cache failed")
				}
				log.Log().Info("CacheKey: ", cacheKey, " refresh cache success")
				return nil, nil
			})
			log.Log().Info("CacheKey: ", cacheKey, " refresh cache shared: ", shared)
		}()
		//结束执行
		return
	}

	// 缓存中没有，直接从k8s中获取
	objList = client.List(config, c.Param("namespace"))
	//在redis中缓存相关查询
	if !KubeClusterService().SaveUnstructuredListToRedis(kubeCluster.Name+":"+c.Request.URL.String(), objList) {
		log.Log().Warn(kubeCluster.Name + ":" + c.Request.URL.String() + "set cache failed")
	}
	c.JSON(200, objList)
}

func (s KubeResourceServiceImpl) IsClusterScoped() bool {
	return s.isClusterScoped
}

func (s KubeResourceServiceImpl) Create(c *gin.Context) {
	// 实现创建逻辑
}

func (s KubeResourceServiceImpl) Update(c *gin.Context) {
	// 实现更新逻辑
}

func (s KubeResourceServiceImpl) Delete(c *gin.Context) {
	// 实现删除逻辑
}

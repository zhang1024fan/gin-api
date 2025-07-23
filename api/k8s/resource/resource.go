package resource

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type DynamicResourceClient struct {
	GVR             schema.GroupVersionResource
	IsClusterScoped bool
}

func (h *DynamicResourceClient) GetResource(c *gin.Context, namespace, name string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (h *DynamicResourceClient) registerCustomRoutes(group *gin.RouterGroup) {
	//TODO implement me
	panic("implement me")
}

func NewDynamicResourceClient(gvr schema.GroupVersionResource, isClusterScoped bool) *DynamicResourceClient {
	return &DynamicResourceClient{
		GVR:             gvr,
		IsClusterScoped: isClusterScoped,
	}
}

func (h *DynamicResourceClient) List(inConfig *rest.Config, namespace string) (objList *unstructured.UnstructuredList) {
	dynamicClient, err := dynamic.NewForConfig(inConfig)
	if err != nil {
		panic(err.Error())
	}
	clientTemp := dynamicClient.Resource(h.GVR)
	if h.IsClusterScoped {
		objList, err = clientTemp.List(context.TODO(), metav1.ListOptions{})
	} else {
		objList, err = dynamicClient.Resource(h.GVR).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	//if err != nil {
	//	result.Failed(c, int(result.ApiCode.FAILED), result.ApiCode.GetMessage(result.ApiCode.FAILED))
	//	return
	//}
	// 在redis中缓存相关查询
	//if !service.KubeClusterService().SaveUnstructuredListToRedis(kubeCluster.Name+":"+c.Request.URL.String(), objList) {
	//	log.Log().Warn(kubeCluster.Name + ":" + c.Request.URL.String() + "set cache failed")
	//}

	//c.JSON(http.StatusOK, objList)
	return

}

func (h *DynamicResourceClient) Get(inConfig *rest.Config, namespace string, name string) (obj *unstructured.Unstructured) {

	dynamicClient, err := dynamic.NewForConfig(inConfig)
	if err != nil {
		panic(err.Error())
	}
	clientTemp := dynamicClient.Resource(h.GVR)

	if h.IsClusterScoped {
		obj, err = clientTemp.Get(context.TODO(), name, metav1.GetOptions{})
	} else {
		obj, err = clientTemp.Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	}

	return
}

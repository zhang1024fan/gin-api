package k8s

import (
	"gin-api/api/k8s/controller"
	"gin-api/api/k8s/service"
	"gin-api/middleware"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func RegisterK8sRoutes(router *gin.RouterGroup) {
	k8sGroup := router.Group("/k8s")
	k8sGroup.GET("/clusters", controller.QueryCluster)
	k8sGroup.POST("/clusters/add", controller.CreateCluster)
	k8sGroup.PUT("/clusters/update", controller.UpdateCluster)

	resourceGroup := k8sGroup.Group("/resource")
	resourceGroup.Use(middleware.Cluster())
	RegisterRoutes(resourceGroup)
}
func RegisterRoutes(group *gin.RouterGroup) {
	handlers := map[string]service.IKubeResourceService{
		"namespaces":             service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("namespaces"), true),
		"deployments":            service.NewKubeResourceService(appsv1.SchemeGroupVersion.WithResource("deployments"), false),
		"pods":                   service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("pods"), false),
		"services":               service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("services"), false),
		"configmaps":             service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("configmaps"), false),
		"nodes":                  service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("nodes"), true),
		"endpoints":              service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("endpoints"), false),
		"endpointslices":         service.NewKubeResourceService(discoveryv1.SchemeGroupVersion.WithResource("endpointslices"), false),
		"secrets":                service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("secrets"), false),
		"persistentvolumes":      service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("persistentvolumes"), true),
		"persistentvolumeclaims": service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("persistentvolumeclaims"), false),
		"serviceaccounts":        service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("serviceaccounts"), false),
		"crds":                   service.NewKubeResourceService(apiextensionsv1.SchemeGroupVersion.WithResource("crds"), true),
		"events":                 service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("events"), false),
		"replicasets":            service.NewKubeResourceService(appsv1.SchemeGroupVersion.WithResource("replicasets"), false),
		"statefulsets":           service.NewKubeResourceService(appsv1.SchemeGroupVersion.WithResource("statefulsets"), false),
		"daemonsets":             service.NewKubeResourceService(appsv1.SchemeGroupVersion.WithResource("daemonsets"), false),
		"ingresses":              service.NewKubeResourceService(networkingv1.SchemeGroupVersion.WithResource("ingresses"), false),
		"jobs":                   service.NewKubeResourceService(batchv1.SchemeGroupVersion.WithResource("jobs"), false),
		"cronjobs":               service.NewKubeResourceService(batchv1.SchemeGroupVersion.WithResource("cronjobs"), false),
		"storageclasses":         service.NewKubeResourceService(storagev1.SchemeGroupVersion.WithResource("storageclasses"), false),
		"roles":                  service.NewKubeResourceService(rbacv1.SchemeGroupVersion.WithResource("roles"), false),
		"rolebindings":           service.NewKubeResourceService(rbacv1.SchemeGroupVersion.WithResource("rolebindings"), false),
		"clusterroles":           service.NewKubeResourceService(rbacv1.SchemeGroupVersion.WithResource("clusterroles"), true),
		"clusterrolebindings":    service.NewKubeResourceService(rbacv1.SchemeGroupVersion.WithResource("clusterrolebindings"), true),
		"podmetrics":             service.NewKubeResourceService(corev1.SchemeGroupVersion.WithResource("metrics.k8s.io"), false),
	}

	for name, handler := range handlers {
		g := group.Group("/" + name)
		if handler.IsClusterScoped() {
			registerClusterScopeRoutes(g, handler)
		} else {
			registerNamespaceScopeRoutes(g, handler)
		}
	}
}

func registerClusterScopeRoutes(group *gin.RouterGroup, handler service.IKubeResourceService) {
	group.GET("", handler.List)
	group.GET("/_all", handler.List)
	group.GET("/_all/:name", handler.Get)
	group.POST("/_all", handler.Create)
	group.PUT("/_all/:name", handler.Update)
	group.DELETE("/_all/:name", handler.Delete)
}

func registerNamespaceScopeRoutes(group *gin.RouterGroup, handler service.IKubeResourceService) {
	group.GET("", handler.List)
	group.GET("/:namespace", handler.List)
	group.GET("/:namespace/:name", handler.Get)
	group.POST("/:namespace", handler.Create)
	group.PUT("/:namespace/:name", handler.Update)
	group.DELETE("/:namespace/:name", handler.Delete)
}

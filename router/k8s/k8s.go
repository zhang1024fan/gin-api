package k8s

import (
	"gin-api/api/k8s/controller"
	"github.com/gin-gonic/gin"
)

func RegisterK8sRoutes(router *gin.RouterGroup) {
	router.GET("/clusters", controller.QueryCluster)
	router.POST("/clusters/add", controller.CreateCluster)
	router.PUT("/clusters/update", controller.UpdateCluster)
}

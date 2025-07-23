package middleware

import (
	"gin-api/api/k8s/service"
	"gin-api/common/result"
	"gin-api/pkg/log"
	"github.com/gin-gonic/gin"
)

const (
	ClusterNameHeader = "x-cluster-name"
	ClusterNameKey    = "cluster-name"
)

// ClusterMiddleware extracts cluster name from header and injects clients into context
func Cluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		clusterName := c.GetHeader(ClusterNameHeader)
		if clusterName == "" {
			if v, ok := c.GetQuery(ClusterNameHeader); ok {
				clusterName = v
			}
		}
		if clusterName == "" {
			result.Failed(c, int(result.ApiCode.KUBEClUSTERNAMENOTEXIST), result.ApiCode.GetMessage(result.ApiCode.KUBEClUSTERNAMENOTEXIST))
			c.Abort()
			return
		}

		log.Log().Info("get ClusterName: ", clusterName)
		//从redis中获取集群信息

		kubecluster := service.KubeClusterService().GetClusterByName(c, clusterName)

		if kubecluster.Name == "" {
			result.Failed(c, int(result.ApiCode.KUBEClUSTERNAMENOTEXIST), result.ApiCode.GetMessage(result.ApiCode.KUBEClUSTERNAMENOTEXIST))
			c.Abort()
			return
		}
		c.Set("cluster", &kubecluster)
		c.Set(ClusterNameKey, kubecluster.Name)
		c.Next()
	}
}

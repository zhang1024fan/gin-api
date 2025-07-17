package model

import (
	"gin-api/common/util"
)

// KubeCluster 表示一条集群记录
type KubeCluster struct {
	ID            int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string     `gorm:"type:varchar(255);not null" json:"name"`                 // 集群名
	KubeConfig    string     `gorm:"column:kubeconfig;type:text;not null" json:"kubeconfig"` // kubeconfig 文件内容
	PrometheusURL string     `gorm:"type:varchar(255)" json:"prometheus_url"`                // Prometheus 监控地址
	IsDefault     bool       `gorm:"type:tinyint(1);default:0;not null" json:"is_default"`   // 是否默认展示
	Version       string     `gorm:"type:varchar(255)" json:"version"`                       // Kubernetes 版本
	CreateTime    util.HTime `gorm:"column:create_time;not null" json:"create_time"`         // 创建时间
	UpdateTime    util.HTime `gorm:"column:update_time;not null" json:"update_time"`         // 更新时间
}

func (KubeCluster) TableName() string {
	return "kube_cluster"
}

// 返回集群列表对象
type KubeClusterVo struct {
	Id        uint   `json:"id"`        // ID
	Name      string `json:"name"`      // 集群名
	IsDefault bool   `json:"isDefault"` //是否默认展示
	Version   string `json:"version"`   //Kubernetes 版本
}

// 创建集群对象
type AddKubeClusterDto struct {
	Name          string `json:"name" binding:"required"` // 集群名
	KubeConfig    string `json:"kubeconfig" binding:"required"`
	PrometheusURL string `json:"prometheus_url" binding:"required"`
	IsDefault     bool   `json:"isDefault"` //是否默认展示
}

// 修改集群对象
type UpdateKubeClusterDto struct {
	ID            int    `json:"id" binding:"required"` // 集群名
	KubeConfig    string `json:"kubeconfig" binding:"required"`
	PrometheusURL string `json:"prometheus_url"`
	IsDefault     bool   `json:"isDefault"` //是否默认展示
}

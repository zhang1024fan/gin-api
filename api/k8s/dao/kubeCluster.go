package dao

import (
	"gin-api/api/k8s/model"
	. "gin-api/pkg/db"
)

// 部门下拉列表
func QueryKubeClusterVoList() (kubeClusterVo []model.KubeClusterVo) {
	Db.Table("kube_cluster").Select("id,name,is_default,version").Scan(&kubeClusterVo)
	return kubeClusterVo
}

// 插入集群信息
func AddKubeCluster(model *model.KubeCluster) bool {
	if err := Db.Create(model).Error; err != nil {
		return false
	}
	return true
}

// 修改集群信息
func UpdateKubeCluster(model *model.KubeCluster) bool {
	if err := Db.Updates(model).Where("id = ?", model.ID).Error; err != nil {
		return false
	}
	return true
}

func GetKubeClusterByName(name string) (kubeCluster model.KubeCluster) {
	Db.Where("name = ?", name).First(&kubeCluster)
	return kubeCluster
}
func GetKubeClusterByID(id int) (kubeCluster model.KubeCluster) {
	Db.Where("id = ?", id).First(&kubeCluster)
	return kubeCluster
}

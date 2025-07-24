package dao

import (
	"gin-api/api/cmdb/model"
	"gin-api/common"

	"gorm.io/gorm"
)

type CmdbHostDao struct {
	db *gorm.DB
}

func NewCmdbHostDao() CmdbHostDao {
	return CmdbHostDao{
		db: common.GetDB(),
	}
}

// 获取主机列表(分页)
func (d *CmdbHostDao) GetCmdbHostListWithPage(page, pageSize int) ([]model.CmdbHost, int64) {
	var list []model.CmdbHost
	var total int64
	
	d.db.Model(&model.CmdbHost{}).Count(&total)
	d.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&list)
	
	return list, total
}

// 获取主机列表
func (d *CmdbHostDao) GetCmdbHostList() []model.CmdbHost {
	var list []model.CmdbHost
	d.db.Find(&list)
	return list
}

// 根据ID获取主机
func (d *CmdbHostDao) GetCmdbHostById(id uint) (model.CmdbHost, error) {
	var host model.CmdbHost
	err := d.db.Where("id = ?", id).First(&host).Error
	return host, err
}

// 根据名称获取主机
func (d *CmdbHostDao) GetCmdbHostByName(name string) (model.CmdbHost, error) {
	var host model.CmdbHost
	err := d.db.Where("host_name = ?", name).First(&host).Error
	return host, err
}

// 检查主机名称是否存在
func (d *CmdbHostDao) CheckNameExists(name string) bool {
	var count int64
	d.db.Model(&model.CmdbHost{}).Where("host_name = ?", name).Count(&count)
	return count > 0
}

// 创建主机
func (d *CmdbHostDao) CreateCmdbHost(host *model.CmdbHost) error {
	return d.db.Create(host).Error
}

// 更新主机
func (d *CmdbHostDao) UpdateCmdbHost(id uint, host *model.CmdbHost) error {
	return d.db.Model(&model.CmdbHost{}).Where("id = ?", id).Updates(host).Error
}

// 删除主机
func (d *CmdbHostDao) DeleteCmdbHost(id uint) error {
	return d.db.Delete(&model.CmdbHost{}, id).Error
}

// 根据分组ID获取主机列表
func (d *CmdbHostDao) GetCmdbHostsByGroupId(groupId uint) []model.CmdbHost {
	var list []model.CmdbHost
	d.db.Where("group_id = ?", groupId).Find(&list)
	return list
}

// 根据主机名称模糊查询
func (d *CmdbHostDao) GetCmdbHostsByHostNameLike(name string) []model.CmdbHost {
	var list []model.CmdbHost
	d.db.Where("host_name LIKE ?", "%"+name+"%").Find(&list)
	return list
}

// 根据IP查询(匹配内网IP、公网IP或SSH IP)
func (d *CmdbHostDao) GetCmdbHostsByIP(ip string) []model.CmdbHost {
	var list []model.CmdbHost
	d.db.Where("private_ip = ? OR public_ip = ? OR ssh_ip = ?", ip, ip, ip).Find(&list)
	return list
}

// 根据状态查询
func (d *CmdbHostDao) GetCmdbHostsByStatus(status int) []model.CmdbHost {
	var list []model.CmdbHost
	d.db.Where("status = ?", status).Find(&list)
	return list
}

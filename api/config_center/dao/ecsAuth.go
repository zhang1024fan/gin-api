package dao

import (
	"gin-api/api/config_center/model"
	"gin-api/common"

	"gorm.io/gorm"
)

type EcsAuthDao struct {
	db *gorm.DB
}

func NewEcsAuthDao() EcsAuthDao {
	return EcsAuthDao{
		db: common.GetDB(),
	}
}

func (d *EcsAuthDao) GetEcsAuthList() []model.EcsAuth {
	var list []model.EcsAuth
	d.db.Find(&list)
	return list
}

func (d *EcsAuthDao) CheckNameExists(name string) bool {
	var count int64
	d.db.Model(&model.EcsAuth{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func (d *EcsAuthDao) GetEcsAuthByName(name string) (model.EcsAuth, error) {
	var auth model.EcsAuth
	err := d.db.Where("name = ?", name).First(&auth).Error
	return auth, err
}

func (d *EcsAuthDao) CreateEcsAuth(auth *model.EcsAuth) error {
	return d.db.Create(auth).Error
}

func (d *EcsAuthDao) UpdateEcsAuth(id uint, auth *model.EcsAuth) error {
	return d.db.Model(&model.EcsAuth{}).Where("id = ?", id).Updates(auth).Error
}

func (d *EcsAuthDao) DeleteEcsAuth(id uint) error {
	return d.db.Delete(&model.EcsAuth{}, id).Error
}

// api/cmdb/dao/cmdbGroup.go

package dao

import (
	"errors"
	"gin-api/api/cmdb/model"
	"gin-api/common/util"
	"gin-api/pkg/db"
	"time"
)

// GetCmdbGroupByNameAndParent 根据名称和父级ID查询是否存在相同分组
func GetCmdbGroupByNameAndParent(name string, parentId uint) (model.CmdbGroup, error) {
	var group model.CmdbGroup
	err := db.Db.Where("name = ? AND parent_id = ?", name, parentId).First(&group).Error
	if err != nil {
		return group, err
	}
	return group, nil
}

// CreateCmdbGroup 创建资产分组（支持根分组和子分组）
func CreateCmdbGroup(group model.CmdbGroup) (bool, error) {
	// 检查是否已有同名 + 同父级的分组
	existGroup, err := GetCmdbGroupByNameAndParent(group.Name, group.ParentID)
	if err == nil && existGroup.ID > 0 {
		return false, errors.New("该父级下已存在同名分组")
	}

	// 设置创建时间
	group.CreateTime = util.HTime{Time: time.Now()}

	// 判断是根分组还是子分组
	if group.ParentID == 0 {
		// 创建根分组
		result := db.Db.Create(&model.CmdbGroup{
			Name:       group.Name,
			ParentID:   0,
			CreateTime: group.CreateTime,
		})
		if result.Error != nil {
			return false, result.Error
		}
	} else {
		// 创建子分组前验证父级是否存在且是根分组
		var parent model.CmdbGroup
		if err := db.Db.Where("id = ? AND parent_id = 0", group.ParentID).First(&parent).Error; err != nil {
			return false, errors.New("父级分组不存在或不是根分组")
		}

		result := db.Db.Create(&model.CmdbGroup{
			Name:       group.Name,
			ParentID:   group.ParentID,
			CreateTime: group.CreateTime,
		})
		if result.Error != nil {
			return false, result.Error
		}
	}

	return true, nil
}

// 查询所有分组
func GetAllGroups() ([]model.CmdbGroup, error) {
	var groups []model.CmdbGroup
	err := db.Db.Order("parent_id asc, id asc").Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// UpdateCmdbGroup 更新资产分组
func UpdateCmdbGroup(group model.CmdbGroup) (bool, error) {
	// 检查分组是否存在
	var existingGroup model.CmdbGroup
	if err := db.Db.First(&existingGroup, group.ID).Error; err != nil {
		return false, errors.New("分组不存在")
	}

	// 检查是否已有同名 + 同父级的分组
	if existingGroup.Name != group.Name || existingGroup.ParentID != group.ParentID {
		existGroup, err := GetCmdbGroupByNameAndParent(group.Name, group.ParentID)
		if err == nil && existGroup.ID > 0 {
			return false, errors.New("该父级下已存在同名分组")
		}
	}

	// 更新分组信息
	result := db.Db.Model(&existingGroup).Updates(model.CmdbGroup{
		Name:     group.Name,
		ParentID: group.ParentID,
	})
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

// DeleteCmdbGroup 删除资产分组
func DeleteCmdbGroup(id uint) (bool, error) {
	// 检查分组是否存在
	var group model.CmdbGroup
	if err := db.Db.First(&group, id).Error; err != nil {
		return false, errors.New("分组不存在")
	}

	// 检查是否有子分组
	var childCount int64
	if err := db.Db.Model(&model.CmdbGroup{}).Where("parent_id = ?", id).Count(&childCount).Error; err != nil {
		return false, err
	}
	if childCount > 0 {
		return false, errors.New("该分组下有子分组，无法删除")
	}

	// 删除分组
	result := db.Db.Delete(&model.CmdbGroup{}, id)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

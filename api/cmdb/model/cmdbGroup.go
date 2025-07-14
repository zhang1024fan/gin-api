// 资产分组
// model/cmdb_group.go
package model

import "gin-api/common/util"

type CmdbGroup struct {
	ID         uint        `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                 // 主键ID
	ParentID   uint        `gorm:"column:parent_id;default:0;comment:'父级分组ID';NOT NULL" json:"parentId"` // 父级分组ID（0 表示根分组）
	Name       string      `gorm:"column:name;varchar(50);comment:'分组名称';NOT NULL" json:"name"`          // 分组名称
	CreateTime util.HTime  `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`         // 创建时间
	Children   []CmdbGroup `json:"children" gorm:"-"`                                                    // 子分组（虚拟字段，用于树形展示）
}

func (CmdbGroup) TableName() string {
	return "cmdb_group"
}

// Id参数
type CmdbGroupIdDto struct {
	Id uint `json:"id"` // ID
}

// BuildTree 构建树形结构
func BuildCmdbGroupTree(groups []CmdbGroup) []CmdbGroup {
	groupMap := make(map[uint]CmdbGroup)
	for _, group := range groups {
		groupMap[group.ID] = group
	}

	var tree []CmdbGroup
	for i := range groups {
		if groups[i].ParentID == 0 {
			tree = append(tree, buildSubTree(groups[i], groupMap))
		}
	}
	return tree
}

// 递归构建子树
func buildSubTree(group CmdbGroup, groupMap map[uint]CmdbGroup) CmdbGroup {
	for _, child := range groupMap {
		if child.ParentID == group.ID {
			group.Children = append(group.Children, buildSubTree(child, groupMap))
		}
	}
	return group
}

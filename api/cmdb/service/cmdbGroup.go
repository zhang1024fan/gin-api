package service

import (
	"gin-api/api/cmdb/dao"
	"gin-api/api/cmdb/model"
	"gin-api/common/constant"
	"gin-api/common/result"
	"gin-api/common/util"
	"time"

	"github.com/gin-gonic/gin"
)

// 接口名称修改为 CmdbGroupServiceInterface 避免冲突
type CmdbGroupServiceInterface interface {
	CreateCmdbGroup(c *gin.Context, group model.CmdbGroup) // 创建分组
	GetAllCmdbGroups(c *gin.Context)                       // 获取所有分组
	UpdateCmdbGroup(c *gin.Context, group model.CmdbGroup) // 更新分组
	DeleteCmdbGroup(c *gin.Context, id uint)               // 删除分组
	GetCmdbGroupByName(c *gin.Context, name string)        // 根据名称查询分组
}

type CmdbGroupServiceImpl struct{}

// 新增分组
func (s CmdbGroupServiceImpl) CreateCmdbGroup(c *gin.Context, group model.CmdbGroup) {
	dao := dao.NewCmdbGroupDao()
	if dao.CheckNameExists(group.Name) {
		result.FailedWithCode(c, constant.GROUP_EXIST, "分组已存在无法创建")
		return
	}
	group.CreateTime = util.HTime{Time: time.Now()}
	err := dao.CreateCmdbGroup(&group)
	if err != nil {
		result.Failed(c, constant.GROUP_EXIST, "创建分组失败")
		return
	}
	result.Success(c, true)
}

// 查询所有分组并返回树形结构
func (s CmdbGroupServiceImpl) GetAllCmdbGroups(c *gin.Context) {
	dao := dao.NewCmdbGroupDao()
	groups := dao.GetCmdbGroupList()
	result.Success(c, model.BuildCmdbGroupTree(groups))
}

// 更新分组
func (s CmdbGroupServiceImpl) UpdateCmdbGroup(c *gin.Context, group model.CmdbGroup) {
	dao := dao.NewCmdbGroupDao()
	err := dao.UpdateCmdbGroup(group.ID, &group)
	if err != nil {
		result.FailedWithCode(c, constant.GROUP_EXIST, err.Error())
		return
	}
	result.Success(c, true)
}

// 删除分组
func (s CmdbGroupServiceImpl) DeleteCmdbGroup(c *gin.Context, id uint) {
	dao := dao.NewCmdbGroupDao()
	err := dao.DeleteCmdbGroup(id)
	if err != nil {
		result.FailedWithCode(c, constant.GROUP_EXIST, err.Error())
		return
	}
	result.Success(c, true)
}

// 根据名称查询分组
func (s CmdbGroupServiceImpl) GetCmdbGroupByName(c *gin.Context, name string) {
	dao := dao.NewCmdbGroupDao()
	group, err := dao.GetCmdbGroupByName(name)
	if err != nil {
		result.Failed(c, constant.GROUP_EXIST, "查询分组失败")
		return
	}
	result.Success(c, group)
}

// 全局服务调用方法
func GetCmdbGroupService() CmdbGroupServiceInterface {
	return &CmdbGroupServiceImpl{}
}

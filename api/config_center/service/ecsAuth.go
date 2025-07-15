package service

import (
	"gin-api/api/config_center/dao"
	"gin-api/api/config_center/model"
	"gin-api/common/constant"
	"gin-api/common/result"
	"gin-api/common/util"
	"time"

	"github.com/gin-gonic/gin"
)

type EcsAuthServiceInterface interface {
	GetEcsAuthList(c *gin.Context)                                              // 获取所有认证信息
	GetEcsAuthByName(c *gin.Context, name string)                               // 根据名称获取认证信息
	CreateEcsAuth(c *gin.Context, dto *model.CreateEcsPasswordAuthDto)          // 创建认证信息
	UpdateEcsAuth(c *gin.Context, id uint, dto *model.CreateEcsPasswordAuthDto) // 更新认证信息
	DeleteEcsAuth(c *gin.Context, id uint)                                      // 删除认证信息
}

type EcsAuthServiceImpl struct {
	dao dao.EcsAuthDao
}

func (s *EcsAuthServiceImpl) GetEcsAuthList(c *gin.Context) {
	list := s.dao.GetEcsAuthList()
	var vos []model.EcsAuthVo
	for _, auth := range list {
		vos = append(vos, model.EcsAuthVo{
			ID:         auth.ID,
			Name:       auth.Name,
			Type:       auth.Type,
			Username:   auth.Username,
			Password:   auth.Password,
			PublicKey:  auth.PublicKey,
			CreateTime: auth.CreateTime,
			Remark:     auth.Remark,
		})
	}
	result.Success(c, vos)
}

func (s *EcsAuthServiceImpl) CreateEcsAuth(c *gin.Context, dto *model.CreateEcsPasswordAuthDto) {
	// 检查名称是否已存在
	if s.dao.CheckNameExists(dto.Name) {
		result.FailedWithCode(c, constant.ECS_AUTH_NAME_EXISTS, "凭据名称已存在")
		return
	}

	auth := model.EcsAuth{
		Name:       dto.Name,
		Type:       1, // 密码认证
		Username:   dto.Username,
		Password:   dto.Password,
		CreateTime: util.HTime{Time: time.Now()},
		Remark:     dto.Remark,
	}
	err := s.dao.CreateEcsAuth(&auth)
	if err != nil {
		result.FailedWithCode(c, constant.ECS_AUTH_CREATE_FAILED, err.Error())
		return
	}
	result.Success(c, true)
}

func (s *EcsAuthServiceImpl) UpdateEcsAuth(c *gin.Context, id uint, dto *model.CreateEcsPasswordAuthDto) {
	auth := model.EcsAuth{
		Name:     dto.Name,
		Username: dto.Username,
		Password: dto.Password,
		Remark:   dto.Remark,
	}
	err := s.dao.UpdateEcsAuth(id, &auth)
	if err != nil {
		result.FailedWithCode(c, constant.ECS_AUTH_UPDATE_FAILED, err.Error())
		return
	}
	result.Success(c, true)
}

func (s *EcsAuthServiceImpl) DeleteEcsAuth(c *gin.Context, id uint) {
	err := s.dao.DeleteEcsAuth(id)
	if err != nil {
		result.FailedWithCode(c, constant.ECS_AUTH_DELETE_FAILED, err.Error())
		return
	}
	result.Success(c, true)
}

func (s *EcsAuthServiceImpl) GetEcsAuthByName(c *gin.Context, name string) {
	auth, err := s.dao.GetEcsAuthByName(name)
	if err != nil {
		result.FailedWithCode(c, constant.ECS_AUTH_NOT_FOUND, "凭据不存在")
		return
	}

	vo := model.EcsAuthVo{
		ID:         auth.ID,
		Name:       auth.Name,
		Type:       auth.Type,
		Username:   auth.Username,
		Password:   auth.Password,
		PublicKey:  auth.PublicKey,
		CreateTime: auth.CreateTime,
		Remark:     auth.Remark,
	}
	result.Success(c, vo)
}

func GetEcsAuthService() EcsAuthServiceInterface {
	return &EcsAuthServiceImpl{
		dao: dao.NewEcsAuthDao(),
	}
}

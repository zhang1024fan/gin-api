package service

import (
	cmdbDao "gin-api/api/cmdb/dao"
	"gin-api/api/cmdb/model"
	"gin-api/common/constant"
	"gin-api/common/result"
	"gin-api/common/util"
	"time"

	"github.com/gin-gonic/gin"
)

type CmdbHostServiceInterface interface {
	GetCmdbHostList(c *gin.Context)                                // 获取主机列表
	GetCmdbHostById(c *gin.Context, id uint)                       // 根据ID获取主机
	GetCmdbHostByName(c *gin.Context, name string)                  // 根据名称获取主机
	CreateCmdbHost(c *gin.Context, dto *model.CreateCmdbHostDto)    // 创建主机
	UpdateCmdbHost(c *gin.Context, id uint, dto *model.UpdateCmdbHostDto) // 更新主机
	DeleteCmdbHost(c *gin.Context, id uint)                         // 删除主机
	GetCmdbHostsByGroupId(c *gin.Context, groupId uint)            // 根据分组ID获取主机列表
}

type CmdbHostServiceImpl struct {
	dao cmdbDao.CmdbHostDao
	groupDao cmdbDao.CmdbGroupDao
}

// 获取主机列表
func (s *CmdbHostServiceImpl) GetCmdbHostList(c *gin.Context) {
	list := s.dao.GetCmdbHostList()
	var vos []model.CmdbHostVo
	for _, host := range list {
		group, _ := s.groupDao.GetCmdbGroupById(host.GroupID)
		vos = append(vos, model.CmdbHostVo{
			ID:          host.ID,
			HostName:    host.HostName,
			GroupID:     host.GroupID,
			GroupName:   group.Name,
			PrivateIP:   host.PrivateIP,
			PublicIP:    host.PublicIP,
			SSHIP:       host.SSHIP,
			SSHName:     host.SSHName,
			SSHKeyID:    host.SSHKeyID,
			SSHPort:     host.SSHPort,
			Remark:      host.Remark,
			Vendor:      host.Vendor,
			Region:      host.Region,
			InstanceID:  host.InstanceID,
			OS:          host.OS,
			Status:      host.Status,
			CPU:         host.CPU,
			Memory:      host.Memory,
			Disk:        host.Disk,
			BillingType: host.BillingType,
			CreateTime:  host.CreateTime,
			ExpireTime:  host.ExpireTime,
			UpdateTime:  host.UpdateTime,
		})
	}
	result.Success(c, vos)
}

// 创建主机
func (s *CmdbHostServiceImpl) CreateCmdbHost(c *gin.Context, dto *model.CreateCmdbHostDto) {
	// 检查名称是否已存在
	if s.dao.CheckNameExists(dto.HostName) {
		result.FailedWithCode(c, constant.CMDB_HOST_NAME_EXISTS, "主机名称已存在")
		return
	}
	// 初始保存连接信息
	host := model.CmdbHost{
		HostName:   dto.HostName,
		GroupID:    dto.GroupID,
		SSHIP:      dto.SSHIP,
		SSHName:    dto.SSHName,
		SSHKeyID:   dto.SSHKeyID,
		SSHPort:    dto.SSHPort,
		Remark:     dto.Remark,
		CreateTime: util.HTime{Time: time.Now()},
		Status:     2, // 初始状态设为未认证
	}
	err := s.dao.CreateCmdbHost(&host)
	if err != nil {
		result.FailedWithCode(c, constant.CMDB_HOST_CREATE_FAILED, err.Error())
		return
	}
	result.Success(c, true)
}

// 更新主机
func (s *CmdbHostServiceImpl) UpdateCmdbHost(c *gin.Context, id uint, dto *model.UpdateCmdbHostDto) {
	// 不再需要查询认证凭证信息，直接从dto获取SSHName和SSHPort

	host := model.CmdbHost{
		HostName:  dto.HostName,
		GroupID:   dto.GroupID,
		SSHIP:     dto.SSHIP,
		SSHName:   dto.SSHName,
		SSHKeyID:  dto.SSHKeyID,
		SSHPort:   dto.SSHPort,
		Remark:    dto.Remark,
	}
	err := s.dao.UpdateCmdbHost(id, &host)
	if err != nil {
		result.FailedWithCode(c, constant.CMDB_HOST_UPDATE_FAILED, err.Error())
		return
	}
	result.Success(c, true)
}

// 根据ID获取主机
func (s *CmdbHostServiceImpl) GetCmdbHostById(c *gin.Context, id uint) {
	host, err := s.dao.GetCmdbHostById(id)
	if err != nil {
		result.FailedWithCode(c, constant.CMDB_HOST_NOT_FOUND, "主机不存在")
		return
	}

	group, _ := s.groupDao.GetCmdbGroupById(host.GroupID)
	vo := model.CmdbHostVo{
		ID:          host.ID,
		HostName:    host.HostName,
		GroupID:     host.GroupID,
		GroupName:   group.Name,
		PrivateIP:   host.PrivateIP,
		PublicIP:    host.PublicIP,
		SSHIP:       host.SSHIP,
		SSHName:     host.SSHName,
		SSHKeyID:    host.SSHKeyID,
		SSHPort:     host.SSHPort,
		Remark:      host.Remark,
		Vendor:      host.Vendor,
		Region:      host.Region,
		InstanceID:  host.InstanceID,
		OS:          host.OS,
		Status:      host.Status,
		CPU:         host.CPU,
		Memory:      host.Memory,
		Disk:        host.Disk,
		BillingType: host.BillingType,
		CreateTime:  host.CreateTime,
		ExpireTime:  host.ExpireTime,
		UpdateTime:  host.UpdateTime,
	}
	result.Success(c, vo)
}

// 根据名称获取主机
func (s *CmdbHostServiceImpl) GetCmdbHostByName(c *gin.Context, name string) {
	host, err := s.dao.GetCmdbHostByName(name)
	if err != nil {
		result.FailedWithCode(c, constant.CMDB_HOST_NOT_FOUND, "主机不存在")
		return
	}

	group, _ := s.groupDao.GetCmdbGroupById(host.GroupID)
	vo := model.CmdbHostVo{
		ID:          host.ID,
		HostName:    host.HostName,
		GroupID:     host.GroupID,
		GroupName:   group.Name,
		PrivateIP:   host.PrivateIP,
		PublicIP:    host.PublicIP,
		SSHIP:       host.SSHIP,
		SSHName:     host.SSHName,
		SSHKeyID:    host.SSHKeyID,
		SSHPort:     host.SSHPort,
		Remark:      host.Remark,
		Vendor:      host.Vendor,
		Region:      host.Region,
		InstanceID:  host.InstanceID,
		OS:          host.OS,
		Status:      host.Status,
		CPU:         host.CPU,
		Memory:      host.Memory,
		Disk:        host.Disk,
		BillingType: host.BillingType,
		CreateTime:  host.CreateTime,
		ExpireTime:  host.ExpireTime,
		UpdateTime:  host.UpdateTime,
	}
	result.Success(c, vo)
}

// 删除主机
func (s *CmdbHostServiceImpl) DeleteCmdbHost(c *gin.Context, id uint) {
	err := s.dao.DeleteCmdbHost(id)
	if err != nil {
		result.FailedWithCode(c, constant.CMDB_HOST_DELETE_FAILED, err.Error())
		return
	}
	result.Success(c, true)
}

// 根据分组ID获取主机列表
func (s *CmdbHostServiceImpl) GetCmdbHostsByGroupId(c *gin.Context, groupId uint) {
	list := s.dao.GetCmdbHostsByGroupId(groupId)
	var vos []model.CmdbHostVo
	for _, host := range list {
		group, _ := s.groupDao.GetCmdbGroupById(host.GroupID)
		vos = append(vos, model.CmdbHostVo{
			ID:          host.ID,
			HostName:    host.HostName,
			GroupID:     host.GroupID,
			GroupName:   group.Name,
			PrivateIP:   host.PrivateIP,
			PublicIP:    host.PublicIP,
			SSHIP:       host.SSHIP,
			SSHName:     host.SSHName,
			SSHKeyID:    host.SSHKeyID,
			SSHPort:     host.SSHPort,
			Remark:      host.Remark,
			Vendor:      host.Vendor,
			Region:      host.Region,
			InstanceID:  host.InstanceID,
			OS:          host.OS,
			Status:      host.Status,
			CPU:         host.CPU,
			Memory:      host.Memory,
			Disk:        host.Disk,
			BillingType: host.BillingType,
			CreateTime:  host.CreateTime,
			ExpireTime:  host.ExpireTime,
			UpdateTime:  host.UpdateTime,
		})
	}
	result.Success(c, vos)
}

func GetCmdbHostService() CmdbHostServiceInterface {
	return &CmdbHostServiceImpl{
		dao: cmdbDao.NewCmdbHostDao(),
		groupDao: cmdbDao.NewCmdbGroupDao(),
	}
}

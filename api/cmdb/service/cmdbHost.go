package service

import (
	"fmt"
	cmdbDao "gin-api/api/cmdb/dao"
	configDao "gin-api/api/config_center/dao"
	"gin-api/api/cmdb/model"
	"gin-api/common/constant"
	"gin-api/common/result"
	"gin-api/common/util"
	"time"

	"github.com/gin-gonic/gin"
)

type CmdbHostServiceInterface interface {
	GetCmdbHostList(c *gin.Context)                                // 获取主机列表
	GetCmdbHostListWithPage(c *gin.Context, page, pageSize int)    // 获取主机列表(分页)
	GetCmdbHostById(c *gin.Context, id uint)                       // 根据ID获取主机
	GetCmdbHostByName(c *gin.Context, name string)                  // 根据名称获取主机
	CreateCmdbHost(c *gin.Context, dto *model.CreateCmdbHostDto)    // 创建主机
	UpdateCmdbHost(c *gin.Context, id uint, dto *model.UpdateCmdbHostDto) // 更新主机
	DeleteCmdbHost(c *gin.Context, id uint)                         // 删除主机
	GetCmdbHostsByGroupId(c *gin.Context, groupId uint)            // 根据分组ID获取主机列表
	GetCmdbHostsByHostNameLike(c *gin.Context, name string)        // 根据主机名称模糊查询
	GetCmdbHostsByIP(c *gin.Context, ip string)                    // 根据IP查询(内网/公网/SSH)
	GetCmdbHostsByStatus(c *gin.Context, status int)              // 根据状态查询
}

type CmdbHostServiceImpl struct {
	dao cmdbDao.CmdbHostDao
	groupDao cmdbDao.CmdbGroupDao
}

// 获取主机列表(分页)
func (s *CmdbHostServiceImpl) GetCmdbHostListWithPage(c *gin.Context, page, pageSize int) {
	list, total := s.dao.GetCmdbHostListWithPage(page, pageSize)
	var vos []model.CmdbHostVo
	for _, host := range list {
		group, _ := s.groupDao.GetCmdbGroupById(host.GroupID)
		vos = append(vos, model.CmdbHostVo{
			ID:          host.ID,
			HostName:    host.HostName,
			Name:        host.Name,
			GroupID:     host.GroupID,
			GroupName:   group.Name,
			PrivateIP:   host.PrivateIP,
			PublicIP:    host.PublicIP,
			SSHIP:       host.SSHIP,
			SSHName:     host.SSHName,
			SSHKeyID:    host.SSHKeyID,
			SSHPort:     host.SSHPort,
			Remark:      host.Remark,
			Vendor:      fmt.Sprintf("%d", host.Vendor),
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
	result.SuccessWithPage(c, vos, total, page, pageSize)
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
			Name:        host.Name,
			GroupID:     host.GroupID,
			GroupName:   group.Name,
			PrivateIP:   host.PrivateIP,
			PublicIP:    host.PublicIP,
			SSHIP:       host.SSHIP,
			SSHName:     host.SSHName,
			SSHKeyID:    host.SSHKeyID,
			SSHPort:     host.SSHPort,
			Remark:      host.Remark,
			Vendor:      fmt.Sprintf("%d", host.Vendor),
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

	// 获取SSH凭据 (前端已确保SSHKeyID有效)
	authDao := configDao.NewEcsAuthDao()
	auth, _ := authDao.GetEcsAuthById(dto.SSHKeyID)

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
		Vendor: 	1,  // 默认创建主机都是为自建主机
		Status:     2, // 初始状态设为未认证
	}

	// 先保存基本信息
	if err := s.dao.CreateCmdbHost(&host); err != nil {
		result.FailedWithCode(c, constant.CMDB_HOST_CREATE_FAILED, err.Error())
		return
	}

	// 立即返回成功响应，后台异步执行SSH操作
	go func() {
		// 准备SSH配置
		sshConfig := util.SSHConfig{
			IP:        dto.SSHIP,
			Port:      dto.SSHPort,
			Type:      auth.Type,
			Username:  dto.SSHName,
			Password:  auth.Password,
			PublicKey: auth.PublicKey,
		}

		// 获取系统信息
		fmt.Println("开始尝试SSH连接获取系统信息...")
		fmt.Printf("SSH配置: %+v\n", sshConfig)
		
		sshUtil := util.NewSSHUtil()
		systemInfo, err := sshUtil.GetSystemInfo(&sshConfig)
		if err != nil {
			fmt.Printf("SSH获取系统信息失败: %v\n", err)
			// 更新状态为认证失败
			s.dao.UpdateCmdbHost(host.ID, &model.CmdbHost{Status: 3})
			return
		}
		
		fmt.Printf("成功获取系统信息: %+v\n", systemInfo)
		
		// 验证必要字段是否存在
		if systemInfo["privateIp"] == "" || systemInfo["os"] == "" {
			fmt.Println("警告: 获取的系统信息不完整")
		}

		// 更新主机信息
		updateData := model.CmdbHost{
			PrivateIP:  systemInfo["privateIp"],
			PublicIP:   systemInfo["publicIp"],
			Name:       systemInfo["name"], // 添加name字段
			OS:         systemInfo["os"],
			CPU:        systemInfo["cpu"],
			Memory:     systemInfo["memory"],
			Disk:       systemInfo["disk"],
			Status:     1, // 认证成功
			UpdateTime: util.HTime{Time: time.Now()},
		}
		s.dao.UpdateCmdbHost(host.ID, &updateData)
	}()

	// 返回成功响应，前端可以通过轮询获取最新状态
	result.Success(c, gin.H{
		"id":     host.ID,
		"status": host.Status,
		"msg":    "主机创建成功，系统信息收集中...",
	})
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
		Name:        host.Name,
		GroupID:     host.GroupID,
		GroupName:   group.Name,
		PrivateIP:   host.PrivateIP,
		PublicIP:    host.PublicIP,
		SSHIP:       host.SSHIP,
		SSHName:     host.SSHName,
		SSHKeyID:    host.SSHKeyID,
		SSHPort:     host.SSHPort,
		Remark:      host.Remark,
		Vendor:      fmt.Sprintf("%d", host.Vendor), 
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
		Name:        host.Name,
		GroupID:     host.GroupID,
		GroupName:   group.Name,
		PrivateIP:   host.PrivateIP,
		PublicIP:    host.PublicIP,
		SSHIP:       host.SSHIP,
		SSHName:     host.SSHName,
		SSHKeyID:    host.SSHKeyID,
		SSHPort:     host.SSHPort,
		Remark:      host.Remark,
		Vendor:      fmt.Sprintf("%d", host.Vendor),
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
			Name:        host.Name,
			GroupID:     host.GroupID,
			GroupName:   group.Name,
			PrivateIP:   host.PrivateIP,
			PublicIP:    host.PublicIP,
			SSHIP:       host.SSHIP,
			SSHName:     host.SSHName,
			SSHKeyID:    host.SSHKeyID,
			SSHPort:     host.SSHPort,
			Remark:      host.Remark,
			Vendor:      fmt.Sprintf("%d", host.Vendor),
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

// 根据主机名称模糊查询
func (s *CmdbHostServiceImpl) GetCmdbHostsByHostNameLike(c *gin.Context, name string) {
	list := s.dao.GetCmdbHostsByHostNameLike(name)
	var vos []model.CmdbHostVo
	for _, host := range list {
		group, _ := s.groupDao.GetCmdbGroupById(host.GroupID)
		vos = append(vos, model.CmdbHostVo{
			ID:          host.ID,
			HostName:    host.HostName,
			Name:        host.Name,
			GroupID:     host.GroupID,
			GroupName:   group.Name,
			PrivateIP:   host.PrivateIP,
			PublicIP:    host.PublicIP,
			SSHIP:       host.SSHIP,
			SSHName:     host.SSHName,
			SSHKeyID:    host.SSHKeyID,
			SSHPort:     host.SSHPort,
			Remark:      host.Remark,
			Vendor:      fmt.Sprintf("%d", host.Vendor),
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

// 根据IP查询(内网/公网/SSH)
func (s *CmdbHostServiceImpl) GetCmdbHostsByIP(c *gin.Context, ip string) {
	list := s.dao.GetCmdbHostsByIP(ip)
	var vos []model.CmdbHostVo
	for _, host := range list {
		group, _ := s.groupDao.GetCmdbGroupById(host.GroupID)
		vos = append(vos, model.CmdbHostVo{
			ID:          host.ID,
			HostName:    host.HostName,
			Name:        host.Name,
			GroupID:     host.GroupID,
			GroupName:   group.Name,
			PrivateIP:   host.PrivateIP,
			PublicIP:    host.PublicIP,
			SSHIP:       host.SSHIP,
			SSHName:     host.SSHName,
			SSHKeyID:    host.SSHKeyID,
			SSHPort:     host.SSHPort,
			Remark:      host.Remark,
			Vendor:      fmt.Sprintf("%d", host.Vendor),
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

// 根据状态查询
func (s *CmdbHostServiceImpl) GetCmdbHostsByStatus(c *gin.Context, status int) {
	list := s.dao.GetCmdbHostsByStatus(status)
	var vos []model.CmdbHostVo
	for _, host := range list {
		group, _ := s.groupDao.GetCmdbGroupById(host.GroupID)
		vos = append(vos, model.CmdbHostVo{
			ID:          host.ID,
			HostName:    host.HostName,
			Name:        host.Name,
			GroupID:     host.GroupID,
			GroupName:   group.Name,
			PrivateIP:   host.PrivateIP,
			PublicIP:    host.PublicIP,
			SSHIP:       host.SSHIP,
			SSHName:     host.SSHName,
			SSHKeyID:    host.SSHKeyID,
			SSHPort:     host.SSHPort,
			Remark:      host.Remark,
			Vendor:      fmt.Sprintf("%d", host.Vendor),
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

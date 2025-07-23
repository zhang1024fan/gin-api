package model

import (
	"gin-api/common/util"
)

// 主机基本信息
type CmdbHost struct {
	ID          uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	HostName    string     `gorm:"column:host_name;varchar(64);comment:'主机名称';NOT NULL" json:"hostName"`
	GroupID     uint       `gorm:"column:group_id;comment:'分组ID';NOT NULL" json:"groupId"`
	PrivateIP   string     `gorm:"column:private_ip;varchar(64);comment:'私网IP'" json:"privateIp"`
	PublicIP    string     `gorm:"column:public_ip;varchar(64);comment:'公网IP'" json:"publicIp"`
	SSHIP       string     `gorm:"column:ssh_ip;varchar(64);comment:'SSH连接IP';NOT NULL" json:"sshIp"`
	SSHName     string     `gorm:"column:ssh_name;varchar(64);comment:'SSH用户名'" json:"sshName"`
	SSHKeyID    uint       `gorm:"column:ssh_key_id;comment:'SSH凭据ID'" json:"sshKeyId"`
	SSHPort     int        `gorm:"column:ssh_port;comment:'SSH端口';default:22" json:"sshPort"`
	Remark      string     `gorm:"column:remark;varchar(500);comment:'备注'" json:"remark"`
	Vendor      string     `gorm:"column:vendor;varchar(32);comment:'厂商(自建/阿里云/腾讯云)'" json:"vendor"`
	Region      string     `gorm:"column:region;varchar(64);comment:'区域'" json:"region"`
	InstanceID  string     `gorm:"column:instance_id;varchar(128);comment:'实例ID'" json:"instanceId"`
	OS          string     `gorm:"column:os;varchar(128);comment:'操作系统'" json:"os"`
	Status      int        `gorm:"column:status;comment:'状态:1->认证成功,2->未认证,3->认证失败'" json:"status"`
	CPU         string     `gorm:"column:cpu;varchar(32);comment:'CPU信息'" json:"cpu"`
	Memory      string     `gorm:"column:memory;varchar(32);comment:'内存信息'" json:"memory"`
	Disk        string     `gorm:"column:disk;varchar(128);comment:'磁盘信息'" json:"disk"`
	BillingType string     `gorm:"column:billing_type;varchar(32);comment:'计费方式'" json:"billingType"`
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
	ExpireTime  util.HTime `gorm:"column:expire_time;comment:'到期时间'" json:"expireTime"`
	UpdateTime  util.HTime `gorm:"column:update_time;comment:'更新时间'" json:"updateTime"`
}

func (CmdbHost) TableName() string {
	return "cmdb_host"
}

// 创建主机DTO - 仅需提供必要连接信息，其他信息将通过SSH连接自动获取
type CreateCmdbHostDto struct {
	HostName  string `validate:"required" json:"hostName"`    // 主机名称(唯一标识)
	GroupID   uint   `validate:"required" json:"groupId"`      // 主机分组ID
	SSHName   string `validate:"required" json:"sshName"`     // SSH登录用户名
	SSHIP     string `validate:"required" json:"sshIp"`         // SSH连接IP(公网或私网IP)
	SSHPort   int    `json:"sshPort"`                          // SSH端口(默认22)
	SSHKeyID  uint   `validate:"required" json:"sshKeyId"`     // SSH凭据ID(从ecsAuth表获取)
	Remark    string `json:"remark"`                            // 备注信息(可选)
}

// 更新主机DTO - 仅需提供必要连接信息
type UpdateCmdbHostDto struct {
	ID         uint   `json:"id"`                              // 主机ID
	HostName   string `validate:"required" json:"hostName"`    // 主机名称(唯一标识)
	GroupID    uint   `validate:"required" json:"groupId"`     // 主机分组ID
	SSHIP      string `validate:"required" json:"sshIp"`        // SSH连接IP(公网或私网IP)
	SSHName    string `validate:"required" json:"sshName"`     // SSH登录用户名
	SSHKeyID   uint   `validate:"required" json:"sshKeyId"`    // SSH凭据ID(从ecsAuth表获取)
	SSHPort    int    `json:"sshPort"`                          // SSH端口(默认22)
	Remark     string `json:"remark"`                           // 备注信息(可选)
}

// 主机ID DTO
type CmdbHostIdDto struct {
	ID uint `json:"id"`
}

// 主机VO
type CmdbHostVo struct {
	ID          uint       `json:"id"`
	HostName    string     `json:"hostName"`
	GroupID     uint       `json:"groupId"`
	GroupName   string     `json:"groupName"`
	PrivateIP   string     `json:"privateIp"`
	PublicIP    string     `json:"publicIp"`
	SSHIP       string     `json:"sshIp"`  // 添加SSH连接IP字段
	SSHName     string     `json:"sshName"`
	SSHKeyID    uint       `json:"sshKeyId"`
	SSHKeyName  string     `json:"sshKeyName"`
	SSHPort     int        `json:"sshPort"`
	Remark      string     `json:"remark"`
	Vendor      string     `json:"vendor"`
	Region      string     `json:"region"`
	InstanceID  string     `json:"instanceId"`
	OS          string     `json:"os"`
	Status      int        `json:"status"`
	CPU         string     `json:"cpu"`
	Memory      string     `json:"memory"`
	Disk        string     `json:"disk"`
	BillingType string     `json:"billingType"`
	CreateTime  util.HTime `json:"createTime"`
	ExpireTime  util.HTime `json:"expireTime"`
	UpdateTime  util.HTime `json:"updateTime"`
}

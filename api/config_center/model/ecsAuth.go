// ECS认证凭证模型
// xiaoRui
package model

import (
	"gin-api/common/util"
)

// ECS认证凭证模型
type EcsAuth struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	Name       string     `gorm:"column:name;varchar(64);comment:'凭证名称';NOT NULL" json:"name"`
	Type       int        `gorm:"column:type;comment:'认证类型:1->密码,2->密钥';NOT NULL" json:"type"`
	Username   string     `gorm:"column:username;varchar(64);comment:'用户名(type=1时使用)'" json:"username"`
	Password   string     `gorm:"column:password;varchar(256);comment:'密码(type=1时使用)'" json:"password"`
	PublicKey  string     `gorm:"column:public_key;type:text;comment:'公钥(type=2时使用)'" json:"publicKey"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
	Remark     string     `gorm:"column:remark;varchar(500);comment:'备注'" json:"remark"`
}

func (EcsAuth) TableName() string {
	return "config_ecsauth"
}

// 创建ECS密码认证DTO
type CreateEcsPasswordAuthDto struct {
	Name     string `validate:"required"` // 凭证名称
	Type     int    `validate:"required"` // 认证类型:1->密码
	Username string `validate:"required"` // 用户名
	Password string `validate:"required"` // 密码
	Remark   string // 备注
}

// 创建ECS密钥认证DTO
type CreateEcsKeyAuthDto struct {
	Name      string `validate:"required"` // 凭证名称
	Type      int    `validate:"required"` // 认证类型:2->密钥
	PublicKey string `validate:"required"` // 公钥
	Username  string `validate:"required"` // 用户名
	Remark    string // 备注
}

// ID参数
type EcsAuthIdDto struct {
	Id uint `json:"id"` // ID
}

// 更新ECS认证DTO
type UpdateEcsAuthDto struct {
	EcsAuthIdDto
	CreateEcsPasswordAuthDto
}

// 认证凭证列表VO
type EcsAuthVo struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Type       int        `json:"type"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	PublicKey  string     `json:"publicKey"`
	CreateTime util.HTime `json:"createTime"`
	Remark     string     `json:"remark"`
}

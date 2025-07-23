// pkg/db/migrate.go
package db

import (
	cmdbmodel "gin-api/api/cmdb/model"
	ccmodel "gin-api/api/config_center/model"

	"gorm.io/gorm"
)

// 注册所有需要自动建表的 model
var models = []interface{}{
	&cmdbmodel.CmdbGroup{},
	&ccmodel.EcsAuth{},
	&cmdbmodel.CmdbHost{},
	// 可以继续添加其他模型...
}

// 自动迁移所有模型
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(models...)
}

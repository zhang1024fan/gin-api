// pkg/db/migrate.go
package db

import (
	"gin-api/api/cmdb/model"
	"gorm.io/gorm"
)

// 注册所有需要自动建表的 model
var models = []interface{}{
	&model.CmdbGroup{},
	// 可以继续添加其他模型...
}

// 自动迁移所有模型
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(models...)
}

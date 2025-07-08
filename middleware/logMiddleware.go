// 操作日志中间件
// author xiaoRui

package middleware

import (
	"gin-api/api/system/dao"
	"gin-api/api/system/model"
	"gin-api/common/util"
	"gin-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.ToLower(c.Request.Method)
		sysAdmin, _ := jwt.GetAdmin(c)
		if method != "get" {
			log := model.SysOperationLog{
				AdminId:    sysAdmin.ID,
				Username:   sysAdmin.Username,
				Method:     method,
				Ip:         c.ClientIP(),
				Url:        c.Request.URL.Path,
				CreateTime: util.HTime{Time: time.Now()},
			}
			dao.CreateSysOperationLog(log)
		}
	}
}

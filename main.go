// 启动程序
// author xiaoRui

package main

import (
	"context"
	"fmt"
	"gin-api/common/config"
	_ "gin-api/docs"
	"gin-api/pkg/db"
	"gin-api/pkg/log"
	"gin-api/pkg/redis"
	"gin-api/router"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// @title 通用后台管理系统
// @version 1.0
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log := log.Log()
	gin.SetMode(config.Config.Server.Model)
	router := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}
	// 启动服务
	go func() {
		log.Info("Conflicting values for 'process.env.NODE_ENV'")
		log.Info("")
		log.Info(fmt.Sprintf("  App running at:"))
		log.Info(fmt.Sprintf("  - Local:   http://%s", config.Config.Server.Address))
		log.Info(fmt.Sprintf("  - Network: http://%s", config.Config.Server.Address))
		log.Info("")
		log.Info("  Note that the development build is not optimized.")
		log.Info("  To create a production build, run npm run build.")
		log.Info("")
		log.Info(fmt.Sprintf("API文档地址: http://%s/swagger/index.html", config.Config.Server.Address))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: %s \n", err)
		}
	}()
	quit := make(chan os.Signal)
	//监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

// 初始化连接
func init() {
	// mysql
	db.SetupDBLink()
	// redis
	redis.SetupRedisDb()
}

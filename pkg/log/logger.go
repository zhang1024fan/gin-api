// pkg/log/logger.go

package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	// 设置默认日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 创建 logs 目录（如果不存在）
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		logger.Errorf("Failed to create logs directory: %v", err)
	}

	// 打开日志文件
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// 同时输出到 stdout 和文件
		mw := io.MultiWriter(os.Stdout, file)
		logger.SetOutput(mw)
	} else {
		logger.Info("Failed to open log file, using stdout only")
	}

	// 🔥 强制设置为 TextFormatter，并验证是否生效
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 - 15:04:05",
	})

	// 输出一条 debug 日志确认当前 formatter 是 TextFormatter
	logger.Debug("Logger initialized with TextFormatter.")
}

// Log 返回全局 Logger 实例，用于业务日志记录
func Log() *logrus.Logger {
	return logger
}

// CustomGinLogger 自定义 Gin 中间件：只输出简洁的文本日志
// pkg/log/logger.go

// CustomGinLogger 自定义 Gin 中间件：输出简洁的文本日志（与 GORM 风格一致）

const (
	bgGreen    = "\x1b[42m"
	bgRed      = "\x1b[41m"
	bgYellow   = "\x1b[43m"
	bgBlue     = "\x1b[44m"
	bgMagenta  = "\x1b[45m"
	bgCyan     = "\x1b[46m"
	colorReset = "\x1b[0m"
)

func getBackgroundColorForStatusCode(code int) string {
	switch {
	case code >= 200 && code < 300:
		return bgGreen
	case code >= 400 && code < 500:
		return bgRed
	case code >= 500:
		return bgYellow
	default:
		return colorReset
	}
}

func getBackgroundColorForMethod(method string) string {
	switch method {
	case "GET":
		return bgBlue
	case "POST":
		return bgMagenta
	case "PUT":
		return bgCyan
	case "DELETE":
		return bgRed
	default:
		return colorReset
	}
}

func CustomGinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()

		// 打印请求开始日志
		logger.WithFields(logrus.Fields{
			"method": method,
			"path":   path,
			"ip":     ip,
		}).Info("Request started")

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		statusBgColor := getBackgroundColorForStatusCode(statusCode)
		methodBgColor := getBackgroundColorForMethod(method)

		// 打印请求完成日志
		logger.WithFields(logrus.Fields{
			"status":  statusCode,
			"latency": latency,
			"method":  method,
			"path":    path,
			"ip":      ip,
		}).Info("Request completed")

		// 保留原有的控制台输出
		fmt.Printf("[GIN] %s | %s%3d%s | %13v | %15s | %s%-6s%s %q\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusBgColor, statusCode, colorReset,
			latency,
			ip,
			methodBgColor, method, colorReset,
			path,
		)
	}
}

// 图片上传 服务层
// author xiaoRui

package service

import (
	"fmt"
	"gin-api/common/config"
	"gin-api/common/result"
	"gin-api/common/util"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"time"
)

type IUploadService interface {
	Upload(c *gin.Context)
}

type UploadServiceImpl struct{}

// 图片上传
func (u UploadServiceImpl) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		result.Failed(c, int(result.ApiCode.FILEUPLOADERROR), result.ApiCode.GetMessage(result.ApiCode.FILEUPLOADERROR))
		return
	}

	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext

	// 文件保存路径（相对于项目根目录）
	filePath := fmt.Sprintf("%s%s/%s",
		config.Config.ImageSettings.UploadDir,
		now.Format("20060102"),
		fileName)

	// 创建目录
	dir := fmt.Sprintf("%s%s",
		config.Config.ImageSettings.UploadDir,
		now.Format("20060102"))

	if err := util.CreateDir(dir); err != nil {
		result.Failed(c, int(result.ApiCode.FILEUPLOADERROR), "创建目录失败: "+err.Error())
		return
	}

	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		result.Failed(c, int(result.ApiCode.FILEUPLOADERROR), "保存文件失败: "+err.Error())
		return
	}

	// 构造浏览器可访问的 URL
	relativeURL := "/upload/" + now.Format("20060102") + "/" + fileName
	finalURL := config.Config.ImageSettings.ImageHost + relativeURL

	// 返回结果
	result.Success(c, finalURL)
}

var uploadService = UploadServiceImpl{}

func UploadService() IUploadService {
	return &uploadService
}

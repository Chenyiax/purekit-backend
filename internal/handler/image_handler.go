package handler

import (
	"net/http"
	"strconv"

	"purekit-backend/errors"
	"purekit-backend/internal/service"
	"purekit-backend/pkg/httputil"

	"github.com/gin-gonic/gin"
)

// ImageHandler 图片处理接口
type ImageHandler struct {
	imageService service.ImageService
}

// NewImageHandler 创建图片处理接口实例
func NewImageHandler(imageService service.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}

// Convert 转换图片格式
// @Summary 转换图片格式
// @Description 将上传的图片转换为指定格式
// @Tags 图片处理
// @Accept multipart/form-data
// @Produce octet-stream
// @Param image formData file true "图片文件"
// @Param format query string true "目标格式 (jpeg, png, gif, webp)"
// @Param quality query int false "图片质量 (1-100)"
// @Success 200 {file} file "转换后的图片"
// @Failure 400 {object} httputil.Response "请求错误"
// @Failure 500 {object} httputil.Response "服务器错误"
// @Router /api/image/convert [post]
func (h *ImageHandler) Convert(c *gin.Context) {
	// 获取目标格式
	format := c.Query("format")
	if format == "" {
		httputil.Error(c, errors.NewBadRequestError("Format parameter is required"))
		return
	}

	// 获取质量参数
	qualityStr := c.DefaultQuery("quality", "85")
	quality, err := strconv.Atoi(qualityStr)
	if err != nil || quality < 1 || quality > 100 {
		quality = 85
	}

	// 获取上传的文件
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		httputil.Error(c, errors.NewBadRequestError("Failed to get image file"))
		return
	}
	defer file.Close()

	// 调用服务进行转换
	result, err := h.imageService.Convert(file, format, quality)
	if err != nil {
		httputil.Error(c, err)
		return
	}

	// 设置响应头
	contentType := "image/" + format
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", "inline; filename=converted."+format)

	// 返回转换后的图片
	c.Data(http.StatusOK, contentType, result)
}

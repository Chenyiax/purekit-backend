package service

import (
	"io"

	"purekit-backend/pkg/imageutil"
)

// ImageServiceImpl 图片转换服务实现
type ImageServiceImpl struct{}

// NewImageService 创建图片转换服务实例
func NewImageService() ImageService {
	return &ImageServiceImpl{}
}

// Convert 转换图片格式
func (s *ImageServiceImpl) Convert(input io.Reader, outputFormat string, quality int) ([]byte, error) {
	// 解码图片
	src, _, err := imageutil.DecodeImage(input)
	if err != nil {
		return nil, err
	}

	// 调用工具包进行转换与编码，内部包含格式校验
	return imageutil.EncodeImage(src, outputFormat, quality)
}

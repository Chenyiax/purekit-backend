package service

import "io"

// ImageService 图片转换服务接口
type ImageService interface {
	// Convert 转换图片格式
	Convert(input io.Reader, outputFormat string, quality int) ([]byte, error)
}

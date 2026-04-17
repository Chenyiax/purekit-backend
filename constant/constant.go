package constant

// 图片格式常量
const (
	ImageFormatJPEG = "jpeg"
	ImageFormatPNG  = "png"
	ImageFormatGIF  = "gif"
	ImageFormatWebP = "webp"
	ImageFormatBMP  = "bmp"
)

// 支持的图片格式列表
var SupportedImageFormats = []string{
	ImageFormatJPEG,
	ImageFormatPNG,
	ImageFormatGIF,
	ImageFormatWebP,
	ImageFormatBMP,
}

// 图片处理常量
const (
	// 默认图片质量
	DefaultImageQuality = 85
	// 最大图片尺寸（5MB）
	MaxImageSize = 5 * 1024 * 1024
	// 最大并发请求数
	MaxConcurrentRequests = 10
	// 请求超时时间（秒）
	RequestTimeout = 30
)

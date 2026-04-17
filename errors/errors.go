package errors

import "errors"

// 错误码定义
const (
	// 系统错误
	ErrInternalServer = "INTERNAL_SERVER_ERROR"
	// 请求错误
	ErrBadRequest = "BAD_REQUEST"
	// 图片处理错误
	ErrImageProcessing = "IMAGE_PROCESSING_ERROR"
	// 不支持的图片格式
	ErrUnsupportedFormat = "UNSUPPORTED_IMAGE_FORMAT"
	// 图片大小超限
	ErrImageTooLarge = "IMAGE_TOO_LARGE"
)

// AppError 自定义业务错误
type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// 错误创建函数
func NewInternalServerError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrInternalServer,
		Message: message,
		Err:     err,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    ErrBadRequest,
		Message: message,
	}
}

func NewImageProcessingError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrImageProcessing,
		Message: message,
		Err:     err,
	}
}

func NewUnsupportedFormatError(message string) *AppError {
	return &AppError{
		Code:    ErrUnsupportedFormat,
		Message: message,
	}
}

func NewImageTooLargeError(message string) *AppError {
	return &AppError{
		Code:    ErrImageTooLarge,
		Message: message,
	}
}

// 检查错误类型
func IsAppError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr)
}

func GetAppError(err error) *AppError {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}
	return nil
}

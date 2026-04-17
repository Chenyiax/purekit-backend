package server

import (
	"fmt"

	"purekit-backend/config"
	"purekit-backend/internal/handler"
	"purekit-backend/internal/middleware"
	"purekit-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// Server HTTP服务器
type Server struct {
	engine *gin.Engine
}

// NewServer 创建HTTP服务器实例
func NewServer() *Server {
	// 创建Gin引擎
	engine := gin.Default()

	// 初始化服务
	imageService := service.NewImageService()
	passwordService := service.NewPasswordService()
	textService := service.NewTextService()
	jsonService := service.NewJsonService()

	// 初始化处理器
	imageHandler := handler.NewImageHandler(imageService)
	passwordHandler := handler.NewPasswordHandler(passwordService)
	textHandler := handler.NewTextHandler(textService)
	jsonHandler := handler.NewJsonHandler(jsonService)

	// 初始化中间件
	rateLimiter := middleware.NewRateLimiter()

	// 注册中间件
	engine.Use(middleware.CORS())
	engine.Use(middleware.SizeLimit())
	engine.Use(middleware.Timeout())
	engine.Use(rateLimiter.Limit())

	// 设置路由
	api := engine.Group("/api")
	{
		image := api.Group("/image")
		{
			image.POST("/convert", imageHandler.Convert)
		}
		password := api.Group("/password")
		{
			password.GET("/generate", passwordHandler.Generate)
		}
		text := api.Group("/text")
		{
			text.POST("/process", textHandler.Process)
		}
		json := api.Group("/json")
		{
			json.POST("/format", jsonHandler.Format)
		}
	}

	return &Server{
		engine: engine,
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", config.AppConfig.Port)
	return s.engine.Run(addr)
}

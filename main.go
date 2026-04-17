package main

import (
	"fmt"
	"log"

	"purekit-backend/config"
	"purekit-backend/server"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化服务器
	srv := server.NewServer()

	// 启动服务器
	fmt.Printf("Server starting on port %d...\n", config.AppConfig.Port)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

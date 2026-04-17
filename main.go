package main

import (
	"io"
	"log"
	"os"

	"purekit-backend/config"
	"purekit-backend/server"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志输出
	logFile, err := os.OpenFile(config.AppConfig.LogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// 设置日志同时输出到文件和控制台
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 初始化服务器
	srv := server.NewServer()

	// 启动服务器
	log.Printf("Server starting on port %d...\n", config.AppConfig.Port)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

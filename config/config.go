package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                  int
	MaxImageSize          int64
	MaxConcurrentRequests int
	RequestTimeout        int
	CorsAllowedOrigins    string
}

var AppConfig *Config

func LoadConfig() error {
	// 加载.env文件
	_ = godotenv.Load()

	var err error
	portStr := getEnv("PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Warning: Invalid PORT %s, using default 8080", portStr)
		port = 8080
	}

	maxSizeStr := getEnv("MAX_IMAGE_SIZE", "5242880")
	maxImageSize, err := strconv.ParseInt(maxSizeStr, 10, 64)
	if err != nil {
		log.Printf("Warning: Invalid MAX_IMAGE_SIZE %s, using default 5MB", maxSizeStr)
		maxImageSize = 5242880
	}

	maxConnStr := getEnv("MAX_CONCURRENT_REQUESTS", "10")
	maxConcurrentRequests, err := strconv.Atoi(maxConnStr)
	if err != nil {
		log.Printf("Warning: Invalid MAX_CONCURRENT_REQUESTS %s, using default 10", maxConnStr)
		maxConcurrentRequests = 10
	}

	timeoutStr := getEnv("REQUEST_TIMEOUT", "30")
	requestTimeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Printf("Warning: Invalid REQUEST_TIMEOUT %s, using default 30", timeoutStr)
		requestTimeout = 30
	}

	AppConfig = &Config{
		Port:                  port,
		MaxImageSize:          maxImageSize,
		MaxConcurrentRequests: maxConcurrentRequests,
		RequestTimeout:        requestTimeout,
		CorsAllowedOrigins:    getEnv("CORS_ALLOWED_ORIGINS", "*"),
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

# PureKit Backend

PureKit 的后端服务，基于 Go 语言和 Gin 框架构建，为前端提供高性能、稳定的工具处理能力。

## 🛠️ 技术栈

- **语言**: Go (Golang)
- **Web 框架**: Gin
- **图片处理**: 标准库 `image` + `golang.org/x/image` (WebP, BMP 支持)
- **配置管理**: `godotenv`
- **开发规范**: 采用 Service-Handler 分层架构

## 📡 API 接口

| 接口路径 | 方法 | 说明 |
| :--- | :--- | :--- |
| `/api/image/convert` | `POST` | 图片格式转换 (支持 query 参数: `format`, `quality`) |
| `/api/password/generate` | `GET` | 生成随机密码 (支持 query 参数: `length`, `upper`, `lower` 等) |
| `/api/text/process` | `POST` | 文本处理 (支持 `upper`, `lower`, `reverse`, `trim`, `collapse`, `cnToEn`, `enToCn`) |
| `/api/json/format` | `POST` | JSON 处理 (支持 `format`, `escape`, `unescape`) |

## 🛡️ 中间件功能

- **CORS**: 全局跨域支持。
- **Size Limit**: 限制上传文件大小（默认 5MB）。
- **Timeout**: 限制请求处理时间（默认 30s）。
- **Rate Limit**: 基础并发限流。

## 🚀 快速开始

### 环境准备
确保已安装 Go 1.20+ 环境。

### 配置文件
复制并根据需要修改 `.env` 文件：
```env
PORT=8080
MAX_IMAGE_SIZE=5242880
REQUEST_TIMEOUT=30
CORS_ALLOWED_ORIGINS=*
```

### 运行服务
```bash
go run main.go
```

## 📂 项目结构

```text
├── config/      # 配置加载
├── constant/    # 常量定义
├── errors/      # 统一错误处理
├── internal/
│   ├── handler/ # 接口逻辑
│   ├── service/ # 核心业务逻辑
│   └── middleware/ # 中间件
├── pkg/
│   ├── httputil/ # HTTP 响应封装
│   └── imageutil/# 图片处理工具
└── server/      # 路由注册与服务器初始化
```
# 构建阶段
FROM golang:1.26-alpine AS builder

# 设置工作目录
WORKDIR /app

# 设置 Go 代理以加快下载（适用于国内环境）
ENV GOPROXY=https://goproxy.cn,direct

# 拷贝依赖描述文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 拷贝源代码
COPY . .

# 编译 Go 程序
# CGO_ENABLED=0 用于生成静态编译的二进制文件，确保在 minimal 镜像中运行
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 运行阶段
FROM alpine:latest

# 安装基础库，如时区数据
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /root/

# 从构建阶段拷贝编译好的二进制文件
COPY --from=builder /app/main .

# 拷贝 .env 文件（可选，取决于你是否使用环境变量替代文件）
COPY --from=builder /app/.env .

# 暴露端口（与 .env 中的 PORT 保持一致）
EXPOSE 8080

# 运行程序
CMD ["./main"]

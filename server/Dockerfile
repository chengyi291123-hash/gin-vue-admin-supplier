# 叭盛供应商管理系统 - 后端 Dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 设置 Go 环境
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码并构建
COPY . .
RUN go build -ldflags="-s -w" -o server .

# 运行阶段
FROM alpine:latest

LABEL MAINTAINER="basheng"

# 设置时区
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata ca-certificates \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

WORKDIR /app

# 复制构建产物
COPY --from=builder /app/server ./
COPY --from=builder /app/config.yaml ./
COPY --from=builder /app/resource ./resource/

EXPOSE 8888

CMD ["./server"]

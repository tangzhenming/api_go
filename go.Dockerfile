# 使用 Go 官方镜像作为基础镜像
FROM golang:1.18 as builder

# 设置工作目录
WORKDIR /app

# 设置 Go 模块代理服务器
ENV GOPROXY=https://goproxy.cn,direct

# 将源代码复制到镜像中
COPY . /app

# 安装依赖
RUN go mod download

# 构建 Go 程序
RUN go build -o backend

# 使用一个轻量级的基础镜像来运行你的程序
FROM alpine:latest

# 将构建好的程序复制到新镜像中
COPY --from=builder /app/backend /backend

# 将 .env 文件复制到新镜像中
COPY --from=builder /app/.env /.env

# 运行你的程序
CMD ["/backend"]

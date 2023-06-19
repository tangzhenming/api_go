# 使用 Go 官方镜像作为基础镜像
FROM golang:1.18 as builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
  GOOS=linux \
  CGO_ENABLED=0 \
  GOARCH=amd64 \
  GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /build

# 将源代码复制到镜像中
COPY . .

# 将我们的代码编译成二进制可执行文件 app
RUN go build -o app .

# 创建一个轻量级的基础镜像
FROM scratch

# 从builder镜像中将构建好的go程序复制到新镜像中
COPY --from=builder /build/app /

# 将 .env 文件复制到新镜像中
COPY --from=builder /build/.env /

# 运行你的程序
ENTRYPOINT ["/app"]

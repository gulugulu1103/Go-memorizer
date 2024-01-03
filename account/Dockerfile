# 使用带有 Go 环境的 Alpine Linux 基础镜像作为构建阶段的基础
FROM golang:alpine as builder

# 设置容器内的工作目录。后续命令将在此目录下执行
WORKDIR /go/src/app

# 开启 Go 模块支持。这适用于 Go 1.11 及更高版本
ENV GO111MODULE=on

# 安装 reflex 工具，用于监控文件变化并自动重新编译 Go 程序
RUN go install github.com/cespare/reflex@latest

# 复制 go.mod 和 go.sum 文件到工作目录，定义了项目的依赖
COPY go.mod .
COPY go.sum .

# 下载并安装依赖。这是基于 go.mod 和 go.sum 文件
RUN go mod download

# 将当前目录下的源代码复制到工作目录
COPY . .

# 编译 Go 程序，并将生成的可执行文件命名为 run
RUN go build -o ./run .

# 开始第二阶段的构建，使用轻量级的 Alpine Linux 基础镜像
FROM alpine:latest

# 安装 CA 证书，允许容器信任 HTTPS 请求中使用的证书
RUN apk --no-cache add ca-certificates

# 设置容器内的工作目录为 /root/
WORKDIR /root/

# 从构建阶段（builder）复制编译好的应用程序到当前工作目录
COPY --from=builder /go/src/app/run .

# 指定容器运行时监听的端口号。应与应用程序的监听端口相匹配
EXPOSE 8080

# 设置容器启动时默认执行的命令，即运行我们的 Go 程序
CMD ["./run"]

# 基于官方的 Golang 镜像进行构建
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将代码复制到容器内的工作目录中
COPY . .

# 设置环境变量
ENV GO111MODULE=on
ENV GOARCH=amd64
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOPATH=/go
ENV GOCACHE=/go/cache
ENV GOPROXY=https://goproxy.cn,direct

# 构建应用程序
RUN go build -o main .

# 运行应用程序
CMD ["./main"]

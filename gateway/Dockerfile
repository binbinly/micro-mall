FROM golang:1.20-alpine as builder

# 作者
LABEL maintainer="mall"

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai

# 执行镜像的工作目录
WORKDIR /go/src/gateway

# 将目录拷贝到镜像里
COPY . .

RUN go mod tidy && go build -o gateway .

# 引入alphine最小linux镜像
FROM alpine

WORKDIR /app
# 复制生成的可执行命令和一些配置文件
COPY --from=builder /go/src/gateway/gateway .
COPY --from=builder /go/src/gateway/configs/docker.yaml .

EXPOSE 9520

# 启动执行命令
ENTRYPOINT ["/app/gateway"]

# 1. build image: docker build -t mall/gateway:latest -f Dockerfile .
# 2. start: docker run --rm -it mall/gateway:latest
# 启动时设置 --rm 选项，这样在容器退出时就能够自动清理容器内部的文件系统
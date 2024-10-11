# 使用官方 Go 作为基础镜像
FROM golang:1.19 as builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=x86_64

# 设置工作目录
WORKDIR /app

COPY ./content-system/go.mod ./content-system/go.sum ./
RUN go mod download
COPY ./content-system .
RUN go build -o /build/content-system


FROM alpine:latest

COPY --from=builder /build/content-system /usr/local/bin/content-system

# 使用更小的基础镜像（如alpine）
FROM alpine:latest

# 暴露服务端口
EXPOSE 9000

# 启动应用
CMD ["/usr/local/bin/content-system"]

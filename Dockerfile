# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.16

# 在容器内创建一个目录来存放我们的应用代码
RUN mkdir /app

# 将工作目录切换到 /app
WORKDIR /app

# 将当前目录下的所有文件拷贝到 /app 目录
COPY . .

# 编译 Go 应用程序
RUN go build -o myapp src/go_code/project_1/main/hello.go

# 暴露 8080 端口
EXPOSE 8080

# 运行应用程序
CMD ["./myapp"]
#源镜像 [国外镜像下载不下来]
FROM golang:latest
#FROM daocloud.io/golang:1.3-onbuild
#作者
MAINTAINER hill "snowspice@163.com"
#设置工作目录
WORKDIR $GOPATH/src/github.com/snowspice/restful-demo
#将服务器的go工程代码加入到docker容器中
# ADD . $GOPATH/src/github.com/snowspice/restful-demo
ADD . .
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 8080

#最终运行docker的命令
ENTRYPOINT  ["./restful-demo"]
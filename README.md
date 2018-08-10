# restful-demo
    基于gin+xorm +log4go 的 restful的demo示例，v1.0
    项目地址：https://github.com/snowspice/restful-demo.git
## 目录结构说明
### common
    存放一些工具类
### config
    配置文件与日志配置文件信息 [当前此文件需要抽离出来,生成可执行文件时，可将cfg.Parse("config/app.properties") 文件路径变为本地绝对路径]
### controller
    控制层，配置路由信息
### service
    服务层，进行数据库操作
### entity
    实体定义层，后续可采用xorm自动生成。
### request
    请求实体对象定义
### main.go
    服务入口文件
## 使用 [windows下]
### 导入项目至 GoLand IDE中【需要先配置好GO运行环境，GO_PATH路径】
    cd $GO_PATH/src
    mkdir github.com
    cd github.com
    mkdir snowspice
    cd snowspice
    git clone https://github.com/snowspice/restful-demo.git
    将工程导入到Goland 中
### 建表语句 user.sql,可自行添加数据
### 运行main.go ，启动服务
    在Goland 的Terminal 执行
     curl -X POST http://127.0.0.1:8080/user/findOne -H "Content-Type:application/x-www-form-urlencoded" -d "userId=1"

### 生成docker 镜像 [由于go环境镜像使用 ：golang:latest ，当前国内网络受限，请翻墙后创建镜像]
    docker build -t restfuldemo_v1.0 .
    创建成功后，会出现：
    Successfully built d5a46802017f
    Successfully tagged restfuldemo_v1.0:latest
    SECURITY WARNING: You are building a Docker image from Windows against a non-Windows Docker host. All files and directories added to build context will have '-rwxr-xr-x' permissions. It is rec
    ommended to double check and reset permissions for sensitive files and directories.

### 运行docker镜像
    docker run --name restfuldemo_v1.0 -p 8080:8080 -d restfuldemo_v1.0
    docker 指令说明：docker --help 查看
    doceker ps|docker images|docker stop|

### 访问启动的镜像服务
    1.查看镜像日志：docker logs -f restfuldemo_v1.0  ：确定服务启动
    2. 浏览器访问：http://localhost:8080/user/1   ：确定服务可用


## go 包管理工具

### govendor [本项目采用govendor管理]
    https://blog.csdn.net/huwh_/article/details/77169858
    https://www.cnblogs.com/hadex/p/6656567.html

### glide
    https://studygolang.com/articles/7129

## 自动生成模型代码 【xorm】
    https://github.com/go-xorm/cmd
    查看帮助信息： xorm help reverse
    示例：【项目根目录下执行一下命令，会在改目录的models下生成文件】
     xorm reverse mysql root:123456@(localhost:3306)/test?charset=utf8  C:/GO_PATH/src/github.com/go-xorm/cmd/xorm/templates/goxorm

## log4go 使用方式
    common.GetLogger().Debug("-----我的测试----")
    需要使用的地方都采用common.GetLogger()














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





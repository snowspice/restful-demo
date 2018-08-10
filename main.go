package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/snowspice/restful-demo/common"
	"net/http"
	"strconv"

	"github.com/snowspice/restful-demo/controller"
)

//将所有方法加入路由
func registerRouter(router *gin.Engine) {

	new(controller.UserController).Router(router)

}

func main() {

	cfg := new(common.Config)
	cfg.Parse("config/app.properties")
	fmt.Println("[ok] load config ")
	common.SetCfg(cfg)

	common.LoggerConfiguration(cfg.Logger["filepath"])

	gin.SetMode(cfg.App["mode"])

	for k, ds := range cfg.Datasource {
		e, err := xorm.NewEngine(ds["driveName"], ds["dataSourceName"])
		fmt.Println("  driveName is s% ,  dataSourceName is s%", ds["driveName"], ds["dataSourceName"])
		if err != nil {
			fmt.Println("data source init error", err.Error())
			return
		}
		fmt.Println("init data source %s", ds["dataSourceName"])
		e.ShowSQL(ds["showSql"] == "true")
		n, _ := strconv.Atoi(ds["maxIdle"])
		e.SetMaxIdleConns(n)
		n, _ = strconv.Atoi(ds["maxOpen"])
		e.SetMaxOpenConns(n)

		common.SetEngin(k, e)
	}
	fmt.Println("[ok] init datasource")
	router := gin.Default()

	registerRouter(router)

	common.Debug("--------------log  debug----------")
	fmt.Println(" [ log debug] is ok")
	err := http.ListenAndServe(cfg.App["addr"]+":"+cfg.App["port"], router)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("[ok] app run", cfg.App["addr"]+":"+cfg.App["port"])
	}
}

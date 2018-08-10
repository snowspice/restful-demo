package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/snowspice/restful-demo/common"
	"github.com/snowspice/restful-demo/service"
	"strconv"

	"encoding/json"
	"fmt"
	"github.com/snowspice/restful-demo/request"
)

type UserController struct {
	common.Controller
}

//用户服务层
var userService service.UserService

//路由注册
func (ctrl *UserController) Router(router *gin.Engine) {

	r := router.Group("user")

	r.POST("findOne", ctrl.findOne)

	r.POST("updateStatus", ctrl.updateStat)

	r.POST("findPage", ctrl.findPage)

	r.GET(":userId",ctrl.findById)

}

// 根据ID 查询指定用户
//curl -X POST http://127.0.0.1:8080/user/findOne -H "Content-Type:application/x-www-form-urlencoded" -d "userId=1"
func (ctrl *UserController) findOne(ctx *gin.Context) {

	userId, _ := strconv.ParseInt(ctx.PostForm("userId"), 10, 64)

	ret := userService.FindOne(userId)
	common.ResultOk(ctx, ret)
}

//get 方法：router  /user/:id
//http://localhost:8080/user/1
func (ctrl *UserController)findById(ctx *gin.Context){
	userId,_:= strconv.ParseInt(ctx.Param("userId"),10,64)
	ret :=userService.FindOne(userId)
	common.ResultOk(ctx,ret)
}

//分页查询数据
//curl -X POST http://127.0.0.1:8080/user/findPage -H "Content-Type:application/x-www-form-urlencoded" -d "{\"page\":1,\"pageSize\":5,\"name\":\"hill\"}"
//curl -X POST http://127.0.0.1:8080/user/findPage -H "Content-Type:application/x-www-form-urlencoded" -d "{\"page\":1,\"pageSize\":5,\"total\":0,\"pageCount\":0,\"nums\":null,\"numsCount\":0,\"name\":\"hill\"}"
//curl -X POST http://127.0.0.1:8080/user/findPage -H "Content-Type:application/json" -d "{\"page\":1,\"pageSize\":5,\"total\":0,\"pageCount\":0,\"nums\":null,\"numsCount\":0,\"name\":\"hill\"}"
func (ctrl *UserController) findPage(ctx *gin.Context) {
	var userArg request.UserArg

	//ctx.ShouldBindWith(&userArg, binding.FormPost)
	ctx.Bind(&userArg)
	//ctx.BindJSON(&userArg)

	if b, err := json.Marshal(userArg); err == nil {
		common.Info("--findPage method request is ->", string(b))
		common.Debug("--findPage method request is ->", string(b))
		fmt.Println("============controller====struct 到json str==")
		fmt.Println(string(b))

	}
	//fmt.Println("---param-->"+string(userArg.Page)+"        name-->"+userArg.Name)

	ret := userService.Query(userArg) //数据
	num := userService.Count(userArg) //总数

	//this.Data["pager"] = tools.CreatePaging(page, pageSize, 365)
	var pager = common.CreatePager(userArg.Page, userArg.PageSize, num)

	common.Info("findPage method response:", ret)
	//最后响应数据列表到前端
	common.ResultList2(ctx, ret, pager)
}

//修改用户状态
//curl -X POST http://127.0.0.1:8080/user/updateStatus -H "Content-Type:application/x-www-form-urlencoded" -d "id=1&stat=0"
func (ctrl *UserController) updateStat(ctx *gin.Context) {

	userId, _ := strconv.ParseInt(ctx.PostForm("id"), 10, 64)
	stat, _ := strconv.Atoi(ctx.PostForm("stat"))
	_, err := userService.UpdateStat(userId, stat)
	if err != nil {
		common.ResultFail(ctx, "修改失败,请稍后再试"+err.Error())
	} else {
		common.ResultOkMsg(ctx, nil, "修改成功,请稍后再试")
	}

}

//注册用户

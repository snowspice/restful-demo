package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/snowspice/restful-demo/common"
	"github.com/snowspice/restful-demo/entity"
	"time"

	"github.com/snowspice/restful-demo/request"

	"fmt"
	"github.com/gin-gonic/gin/json"
)

type UserService struct{}

//【查询】根据userId 获取用户编号
func (service *UserService) FindOne(userId int64) entity.User {
	var user entity.User
	orm := common.OrmEngin()
	orm.Id(userId).Get(&user)
	return user
}

//【修改】更新指定用户的状态
func (service *UserService) UpdateStat(id int64, stat int) (int64, error) {
	var user entity.User
	user.ID = id
	user.Status = stat
	orm := common.OrmEngin()
	r, e := orm.ID(id).Cols("status").Update(&user)
	return r, e
}

//【删除】指定用户
func (service *UserService) Del(userId int64) (int64, error) {
	var user entity.User
	user.ID = userId
	orm := common.OrmEngin()
	r, e := orm.Delete(&user)
	return r, e
}

//【增加】用户
func (service *UserService) Register(ctx *gin.Context, user *entity.User) (p *entity.User, err error) {
	var u entity.User
	orm := common.OrmEngin()
	t := orm.Where("id>0")
	t.Where("name=?", user.Name)
	t.Get(&u)
	if u.ID > 0 {
		err = errors.New("该账户已存在")
		return
	}

	user.Status = 1
	user.Created = common.JsonDateTime(time.Now())
	user.ID, err = orm.InsertOne(user)
	p = user
	return
}

//【分页查询】返回当前页数据
func (service *UserService) Query(arg request.UserArg) []entity.User {
	var users []entity.User = make([]entity.User, 0)

	if arg.Page < 1 {
		arg.Page = 1
	}
	if arg.PageSize < 1 {
		arg.PageSize = 10
	}
	orm := common.OrmEngin()
	t := orm.Where("id>0")
	fmt.Println("====参数 name =====>" + arg.Name)
	if 0 < len(arg.Name) {
		t = t.Where("name like ?", "%"+arg.Name+"%")
	}
	if b, err := json.Marshal(arg); err == nil {

		fmt.Println("================struct 到json str==")
		fmt.Println(string(b))

	}
	t.Limit(int(arg.PageSize), ((int(arg.Page))-1)*(int(arg.Page))).Find(&users)
	return users
}

//获取数量
func (service *UserService) Count(arg request.UserArg) (n int64) {
	var user entity.User
	orm := common.OrmEngin()
	t := orm.Where("id>0")
	if 0 < len(arg.Name) {
		t = t.Where("name like ?", "%"+arg.Name+"%")
	}

	n, _ = t.Count(&user)
	return
}

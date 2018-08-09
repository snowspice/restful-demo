package request

import "github.com/snowspice/restful-demo/common"

type UserArg struct {

	common.Pager
	Name string `form:"name" json:"name"` //自定义字段
}




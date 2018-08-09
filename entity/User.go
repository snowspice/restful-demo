package entity

import (
	"com.mydadao.com/restful-demo/common"
)

type User struct {

	ID int64 `xorm:"pk autoincr 'id'" form:"id" json:"id"`
	Name string `xorm:"varchar(40)" form:"name" json:"name"`
	Status int `xorm:"status" json:"status"`
	Created common.JsonDateTime `xorm:"created" form:"created" json:"created"  time_format:"2006-01-02 15:04:05"`

}
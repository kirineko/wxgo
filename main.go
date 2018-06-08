package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "wxgo/models"
	_ "wxgo/routers"
	"wxgo/utils"
)

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	err := utils.GetAndUpdateDBWxAToken(o)
	if err != nil {
		//todo 向微信请求access_token失败 结合业务逻辑处理
		fmt.Println("get access_token task failed")
	}
	beego.Run()
}

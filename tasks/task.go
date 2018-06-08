package tasks

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"wxgo/utils"
)

func init() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	shopTimeTask(o)
}

func shopTimeTask(o orm.Ormer) {

	timeStr2 := "0 */60 * * * *" // 获取wx token
	t2 := toolbox.NewTask("getAtoken", timeStr2, func() error {

		err := utils.GetAndUpdateDBWxAToken(o)
		if err != nil {
			//todo 向微信请求access_token失败 结合业务逻辑处理
			fmt.Println("get access_token task failed")
		}
		return nil
	})
	toolbox.AddTask("tk2", t2)
	toolbox.StartTask()
}

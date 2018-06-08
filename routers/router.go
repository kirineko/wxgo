package routers

import (
	"github.com/astaxie/beego"
	"wxgo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wx", &controllers.WxController{})
}

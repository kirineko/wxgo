package main

import (
	"github.com/astaxie/beego"
	_ "wxgo/models"
	_ "wxgo/routers"
	_ "wxgo/tasks"
)

func main() {
	beego.Run()
}

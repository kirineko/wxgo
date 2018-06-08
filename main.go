package main

import (
	"github.com/astaxie/beego"
	"wxgo/models"
	_ "wxgo/routers"
)

func main() {
	beego.Run()

	models.Init()
}

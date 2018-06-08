package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type WxAccessToken struct {
	Id          int `orm:"auto"`
	AccessToken string
}

type WxBase struct {
	Id             int `orm:"auto"`
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}

type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
}

func Init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root@/wechat?charset=utf8", 30)
	//注册定义的model
	orm.RegisterModel(new(WxAccessToken), new(WxBase))

	// 创建table
	orm.RunSyncdb("default", false, false)
}

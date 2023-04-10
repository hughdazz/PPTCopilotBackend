package main

import (
	_ "backend/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp(host.docker.internal:3307)/now_db?charset=utf8&loc=Local")
}

func main() {
	beego.Run()
}

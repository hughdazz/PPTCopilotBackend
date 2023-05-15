package main

import (
	_ "backend/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	logs.SetLogger("console")
	beego.Run()
}

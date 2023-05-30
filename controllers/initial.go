package controllers

import (
	"backend/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type InitialController struct {
	beego.Controller
}

func (this *InitialController) Get() {
	// 初始化数据库
	// 删除所有表
	models.DropAllTables()

	// 如果表不存在则创建表
	orm.RunSyncdb("default", false, true)

	// 初始化User
	models.CreateUser("admin", "123456", "admin@163.com")
	models.CreateUser("jack", "123456", "jack@qq.com")
	models.CreateUser("tom", "123456", "tom@gmail.com")
	models.CreateUser("lucy", "123456", "2052150@tongji.edu.cn")

	// 初始化Project
	models.CreateProject("project1", "admin", 1)

	// 初始化File
	models.CreateFile("cover.png", 1)

	this.Data["json"] = MakeResponse(OK, "success", nil)
	this.ServeJSON()
}

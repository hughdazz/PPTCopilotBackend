package controllers

import (
	"backend/models"

	beego "github.com/beego/beego/v2/server/web"
)

type InitController struct {
	beego.Controller
}

func (c *InitController) Get() {
	// 初始化数据库
	// 删除所有表
	models.DropAllTables()

	// 初始化User
	models.CreateUser("admin", "123456", "admin@163.com")
	models.CreateUser("jack", "123456", "jack@qq.com")
	models.CreateUser("tom", "123456", "tom@gmail.com")

}

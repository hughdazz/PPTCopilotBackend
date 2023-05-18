package models

import "github.com/beego/beego/v2/client/orm"

func DropAllTables() {
	o := orm.NewOrm()
	// 删除User表
	o.Raw("DROP TABLE IF EXISTS user").Exec()
	// 删除Project表
	o.Raw("DROP TABLE IF EXISTS project").Exec()
	// 删除File表
	o.Raw("DROP TABLE IF EXISTS file").Exec()
	// 删除Outline表
	o.Raw("DROP TABLE IF EXISTS outline").Exec()
	// 删除Favorite表
	o.Raw("DROP TABLE IF EXISTS favorite").Exec()
	// 删除Template表
	o.Raw("DROP TABLE IF EXISTS template").Exec()
}

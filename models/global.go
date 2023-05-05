package models

import "github.com/beego/beego/v2/client/orm"

func DropAllTables() {
	// 删除所有表数据
	o := orm.NewOrm()
	// 删除User表数据
	o.Raw("delete from user").Exec()
}

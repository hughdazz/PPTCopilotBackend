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
	// 删除GptInfo表
	o.Raw("DROP TABLE IF EXISTS gpt_info").Exec()
	// 删除EmailVerification表
	o.Raw("DROP TABLE IF EXISTS email_verification").Exec()

}

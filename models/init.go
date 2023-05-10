package models

import "github.com/beego/beego/v2/client/orm"

// 初始化数据表
func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp(host.docker.internal:3307)/now_db?charset=utf8&loc=Local")
	// 注册定义的model
	orm.RegisterModel(new(Outline))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(File))

	// 如果表不存在则创建表
	orm.RunSyncdb("default", false, true)
}

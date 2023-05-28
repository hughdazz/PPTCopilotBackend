package models

import (
	"github.com/beego/beego/v2/client/orm"
	"os"
	"fmt"
)

// 初始化数据表
func init() {
	// 获取环境变量
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	fmt.Println("MYSQL_HOST: ", mysqlHost)
	fmt.Println("MYSQL_PORT: ", mysqlPort)
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp("+mysqlHost+":"+mysqlPort+")/now_db?charset=utf8&loc=Local")

	// 注册定义的model
	orm.RegisterModel(new(Outline))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(File))
	orm.RegisterModel(new(Favorite))
	orm.RegisterModel(new(Template))

	// 如果表不存在则创建表
	orm.RunSyncdb("default", false, true)
}

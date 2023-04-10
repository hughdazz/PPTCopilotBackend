package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Project struct {
	Id          int
	Name        string    `orm:"size(100)"`
	Description string    `orm:"size(100)"`
	Creator     *User     `orm:"rel(fk)"` // 设置一对多的反向关系
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

// 初始化数据表
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp(host.docker.internal:3307)/now_db?charset=utf8&loc=Local")

	// 注册定义的model
	orm.RegisterModel(new(Project))
}

func NewProject(name string, description string, creator *User) (Project, error) {
	o := orm.NewOrm()
	project := Project{Name: name, Description: description, Creator: creator}
	_, err := o.Insert(&project)
	return project, err
}

func GetProject(id int) (Project, error) {
	o := orm.NewOrm()
	project := Project{Id: id}
	err := o.Read(&project)
	return project, err
}

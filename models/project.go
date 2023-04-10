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

func NewProject(name string, description string, creator_id int) (Project, error) {
	o := orm.NewOrm()
	var creator User
	creator.Id = creator_id
	err := o.Read(&creator)
	if err != nil {
		return Project{}, err
	}
	project := Project{Name: name, Description: description, Creator: &creator}
	_, err = o.Insert(&project)
	return project, err
}

func GetProject(id int) (Project, error) {
	o := orm.NewOrm()
	project := Project{Id: id}
	err := o.Read(&project)
	return project, err
}

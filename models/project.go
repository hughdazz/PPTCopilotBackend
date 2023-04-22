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
	// 根据id获取用户信息
	err := o.Read(&creator)
	if err != nil {
		// 用户不存在
		return Project{}, err
	}

	project := Project{Name: name, Description: description, Creator: &creator}
	// 创建项目
	_, err = o.Insert(&project)
	return project, err
}

func UpdataProject(id int) error {
	o := orm.NewOrm()
	project := Project{Id: id}
	if err := o.Read(&project); err != nil {
		return err
	}
	if _, err := o.Update(&project); err != nil {
		return err
	}
	return nil
}

func GetProject(id int) (Project, error) {
	o := orm.NewOrm()
	project := Project{Id: id}
	err := o.Read(&project)
	return project, err
}

func GetProjects(id int) ([]Project, error) {
	o := orm.NewOrm()
	var projects []Project
	_, err := o.QueryTable("project").Filter("creator_id", id).All(&projects)
	return projects, err
}

func DeleteProject(id int) error {
	o := orm.NewOrm()
	project := Project{Id: id}
	err := o.Read(&project)
	if err == nil {
		// 使用 Delete() 函数删除对象
		DeleteDir(id)
		if err != nil {
			return err
		}
		_, err := o.Delete(&project)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

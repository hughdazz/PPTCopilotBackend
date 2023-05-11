package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Project struct {
	Id          int       `orm:"auto;pk"`
	Name        string    `orm:"size(100)"`
	Description string    `orm:"size(100)"`
	Creator     *User     `orm:"rel(fk)"` // 设置一对多的反向关系
	Star        int       `orm:"default(0)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

type Favorite struct {
	Id      int       `orm:"auto;pk"`
	User    *User     `orm:"rel(fk)"` // 设置一对多的反向关系
	Project *Project  `orm:"rel(fk)"` // 设置一对多的反向关系
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func RefactProjects(projects []Project) []Project {
	for i, project := range projects {
		creator_temp, _ := GetUser(project.Creator.Id)
		creator := User{Id: creator_temp.Id, Username: creator_temp.Username, Email: creator_temp.Email}
		// projects[i] = Project{Name: project.Name, Description: project.Description, Creator: &creator}
		projects[i].Name = project.Name
		projects[i].Description = project.Description
		projects[i].Creator = &creator
	}
	return projects
}

func RefactProject(project Project) Project {
	creator_temp, _ := GetUser(project.Creator.Id)
	creator := User{Id: creator_temp.Id, Username: creator_temp.Username, Email: creator_temp.Email}
	project.Creator = &creator

	return project
}

func CreateProject(name string, description string, creator_id int) (Project, error) {
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

func UpdateProjectName(id int, name string) (Project, error) {
	o := orm.NewOrm()
	project := Project{Id: id}
	err := o.Read(&project)
	if err == nil {
		project.Name = name
		_, err := o.Update(&project)
		return project, err
	}
	return project, err
}

func UpdateProjectDescription(id int, description string) (Project, error) {
	o := orm.NewOrm()
	project := Project{Id: id}
	err := o.Read(&project)
	if err == nil {
		project.Description = description
		_, err := o.Update(&project)
		return project, err
	}
	return project, err
}

func IncProjectStar(id int) (Project, error) {
	o := orm.NewOrm()
	project := Project{Id: id}
	err := o.Read(&project)
	if err == nil {
		project.Star = project.Star + 1
		_, err := o.Update(&project)
		return project, err
	}
	return project, err
}

func GetProject(id int) (Project, error) {
	o := orm.NewOrm()
	// 查找，使用
	var project Project
	err := o.QueryTable("project").Filter("id", id).One(&project)
	return project, err
}

func GetProjects(id int) ([]Project, error) {
	o := orm.NewOrm()
	var projects []Project
	_, err := o.QueryTable("project").Filter("creator_id", id).All(&projects)
	return projects, err
}

func GetAllProjects() []Project {
	o := orm.NewOrm()
	var projects []Project
	o.QueryTable("project").All(&projects)
	return projects
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

func SearchProjects(keywords []string) ([]Project, error) {
	o := orm.NewOrm()

	// 构造 SQL 语句
	sql := "SELECT * FROM project WHERE 1=1"
	for _, keyword := range keywords {
		// 使用 OR 连接多个关键词
		sql += fmt.Sprintf(" AND (name LIKE '%%%s%%' OR description LIKE '%%%s%%')", keyword, keyword)
	}

	// 执行 SQL 查询
	var projects []Project
	_, err := o.Raw(sql).QueryRows(&projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func LikeProject(user_id int, project_id int) (Favorite, error) {
	o := orm.NewOrm()
	var user User
	user.Id = user_id
	var project Project
	project.Id = project_id
	favorite := Favorite{User: &user, Project: &project}
	_, err := o.Insert(&favorite)
	IncProjectStar(project_id)
	return favorite, err
}

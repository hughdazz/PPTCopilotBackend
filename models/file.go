package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type File struct {
	Id   int
	Name string `orm:"size(100)"`
	// Description string    `orm:"size(100)"`
	Project *Project  `orm:"rel(fk)"` // 设置一对多的反向关系
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func NewFile(name string, project_id int) (File, error) {
	o := orm.NewOrm()
	var project Project
	project.Id = project_id
	// 根据id获取项目信息
	err := o.Read(&project)
	if err != nil {
		// 项目不存在
		return File{}, err
	}

	file := File{Name: name, Project: &project}
	// 创建文件结构
	_, err = o.Insert(&file)
	return file, err
}

func GetFileInProjectByName(file_name string, project_id int) (File, int) {
	//找到该项目下的所有file
	files, err := GetFiles(project_id)
	if err == nil {
		//遍历所有file，找有无重名
		for _, file := range files {
			if file_name == file.Name {
				return file, 0 //找到
			}
		}
		return File{}, 1 //未找到
	}
	return File{}, 2 //查找出错
}

func GetFiles(id int) ([]File, error) {
	o := orm.NewOrm()
	var files []File
	_, err := o.QueryTable("file").Filter("project_id", id).All(&files)
	return files, err
}

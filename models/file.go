package models

import (
	"os"
	"path/filepath"
	"strconv"
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

func DeleteDir(project_id int) error {
	saveDir := "static/files/" + strconv.Itoa(project_id)
	err := os.RemoveAll(saveDir)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(file_name string, project_id int) int {
	file, res := GetFileInProjectByName(file_name, project_id)
	if res != 0 {
		// 文件不存在
		return 1
	}
	o := orm.NewOrm()
	_, err := o.Delete(&file)
	if err != nil {
		return 2 //文件结构删除错误
	}

	path := GetFilePathByName(file_name, project_id)
	err = DeleteFileByPath(path)

	if err != nil {
		//文件删除错误
		return 3
	}
	return 0
}

func DeleteFileByPath(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	err = os.Remove(absPath)
	return err
}

func GetFilePathByName(file_name string, project_id int) string {

	saveDir := "static/files/" + strconv.Itoa(project_id)

	filePath := saveDir + "/" + file_name
	return filePath
}

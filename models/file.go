package models

import (
	"errors"
	"io"
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

func RefactFiles(files []File) []File {
	for i, file := range files {
		project_temp, _ := GetProject(file.Project.Id)
		creator_temp, _ := GetUser(project_temp.Creator.Id)
		creator := User{Id: creator_temp.Id, Username: creator_temp.Username, Email: creator_temp.Email}
		project := Project{Id: project_temp.Id, Name: project_temp.Name, Creator: &creator, Star: project_temp.Star, Description: project_temp.Description, Created: project_temp.Created, Updated: project_temp.Updated}
		files[i].Name = file.Name
		files[i].Project = &project
	}
	return files
}

func GetFile(file_name string, id int) (File, error) {
	//找到该项目下的所有file
	files, err := GetFiles(id)

	if err == nil {
		//遍历所有file，找有无重名
		for _, file := range files {
			if file_name == file.Name {
				return file, nil
			}
		}
		return File{}, errors.New("文件不存在")
	}
	return File{}, err
}

func GetFiles(id int) ([]File, error) {
	o := orm.NewOrm()
	var files []File
	_, err := o.QueryTable("file").Filter("project_id", id).All(&files)
	return files, err
}

func CreateFile(name string, project_id int) (File, error) {
	o := orm.NewOrm()
	var project Project
	project.Id = project_id
	// 根据id获取项目信息
	err := o.Read(&project)
	if err != nil {
		// 项目不存在
		return File{}, err
	}

	_file, err := GetFile(name, project_id)
	if _file.Name == name {
		return File{}, errors.New("文件已存在")
	}

	file := File{Name: name, Project: &project}
	// 创建文件结构
	_, err = o.Insert(&file)
	return file, err
}

func DeleteDir(project_id int) error {
	saveDir := "static/files/" + strconv.Itoa(project_id)
	err := os.RemoveAll(saveDir)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(file_name string, project_id int) error {
	file, err := GetFile(file_name, project_id)
	if err != nil {
		// 文件不存在
		return err
	}
	o := orm.NewOrm()
	_, err = o.Delete(&file)
	if err != nil {
		return err //文件结构删除错误
	}

	path := GetFilePathByName(file_name, project_id)
	err = DeleteFileByPath(path)

	if err != nil {
		//文件删除错误
		return err
	}
	return nil
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

	saveDir := "static/project/" + strconv.Itoa(project_id)

	filePath := saveDir + "/" + file_name
	return filePath
}
func GetSaveDir(project_id int) string {
	saveDir := "static/project/" + strconv.Itoa(project_id)
	return saveDir
}

func CopyFile(srcFile, dstFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

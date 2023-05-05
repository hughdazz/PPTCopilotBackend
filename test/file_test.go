package controllers

import (
	"backend/models"
	"testing"
)

// func GetFilePathByName(file_name string, project_id int) string {

// 	saveDir := "static/files/" + strconv.Itoa(project_id)

// 	filePath := saveDir + "/" + file_name
// 	return filePath
// }

// 如上函数的单元测试函数
func TestGetFilePathByName(t *testing.T) {
	filePath := models.GetFilePathByName("test", 1)
	if filePath != "static/files/1/test" {
		t.Errorf("GetFilePathByName() = %s; want static/files/1/test", filePath)
	}
}

// func GetUser(id int) (User, error) {
// 	o := orm.NewOrm()
// 	user := User{Id: id}
// 	err := o.Read(&user)
// 	return user, err
// }

//INSERT INTO `user` VALUES (1, 'hughdazz', '123456', '1@qq.com');

// 如上函数的单元测试函数
// func TestGetUser(t *testing.T) {
// 	user, err := models.GetUser(1)
// 	if err != nil {
// 		t.Errorf("GetUser() = %s; want nil", err)
// 	}
// 	if user.Username != "hughdazz" {
// 		t.Errorf("GetUser() = %s; want hughdazz", user.Username)
// 	}
// }

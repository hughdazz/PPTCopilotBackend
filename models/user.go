package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// 用户信息
type User struct {
	Id       int
	Username string `orm:"size(100)"`
	Password string `orm:"size(100)"`
	Email    string `orm:"size(100)"`
}

// 验证用户信息
func VerifyUser(username_or_email string, password string) (User, error) {
	// 通过用户名或邮箱获取用户信息
	o := orm.NewOrm()
	user := User{Username: username_or_email}
	err := o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		user = User{Email: username_or_email}
		err = o.Read(&user, "Email")
		if err == orm.ErrNoRows {
			return user, err
		}
	}
	// 验证密码是否正确
	if user.Password != password {
		err = orm.ErrNoRows
	}
	return user, err
}

// 新建用户
func CreateUser(username string, password string, email string) (User, error) {
	o := orm.NewOrm()
	user := User{Username: username, Password: password, Email: email}
	_, err := o.Insert(&user)
	return user, err
}

// 获取用户
func GetUser(id int) (User, error) {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	return user, err
}

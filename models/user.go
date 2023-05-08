package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// 用户信息
type User struct {
	Id          int       `orm:"column(id);pk;auto"`
	Username    string    `orm:"size(100);unique"`
	Password    string    `orm:"size(100);"`
	Email       string    `orm:"size(100);unique"`
	Description string    `orm:"size(100);"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func GetAllUsers() []User {
	o := orm.NewOrm()
	var users []User
	o.QueryTable("user").All(&users)
	return users
}

func UpdateUserUsername(id int, username string) error {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return err
	}
	user.Username = username
	_, err = o.Update(&user, "Username")
	if err != nil {
		return err
	}
	return nil
}
func UpdateUserPassword(id int, password string) error {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return err
	}
	user.Password = password
	_, err = o.Update(&user, "Password")
	if err != nil {
		return err
	}
	return nil
}
func UpdateUserEmail(id int, email string) error {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return err
	}
	user.Email = email
	_, err = o.Update(&user, "Email")
	if err != nil {
		return err
	}
	return nil
}
func UpdateUserDescription(id int, description string) error {
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return err
	}
	user.Description = description
	_, err = o.Update(&user, "Description")
	if err != nil {
		return err
	}
	return nil
}
func DeleteUser(id int) error {
	o := orm.NewOrm()
	user := User{Id: id}
	_, err := o.Delete(&user)
	return err
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

// 更新密码
func UpdatePassword(email string, newPassword string) error {
	// 查询邮箱对应的用户
	o := orm.NewOrm()
	user := User{Email: email}
	err := o.Read(&user, "Email")

	if err != nil {
		return err
	}

	// 更新用户密码
	user.Password = newPassword
	_, err = o.Update(&user, "Password")
	if err != nil {
		return err
	}

	return nil
}

// 通过邮箱获得用户
func GetUserByEmail(email string) (User, error) {
	o := orm.NewOrm()
	user := User{Email: email}
	err := o.Read(&user, "Email")
	return user, err
}

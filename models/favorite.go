package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Favorite struct {
	Id      int       `orm:"auto;pk"`
	User    *User     `orm:"rel(fk)"` // 设置一对多的反向关系
	Project *Project  `orm:"rel(fk)"` // 设置一对多的反向关系
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetFavorites(id int) ([]Favorite, error) {
	o := orm.NewOrm()
	var favorites []Favorite
	_, err := o.QueryTable("favorite").Filter("user_id", id).All(&favorites)
	return favorites, err
}

func RefactFavorites(favorites []Favorite) []Favorite {
	for i, favorite := range favorites {
		project_temp, _ := GetProject(favorite.Project.Id)
		creator_temp, _ := GetUser(project_temp.Creator.Id)
		creator := User{Id: creator_temp.Id, Username: creator_temp.Username, Email: creator_temp.Email}
		project := Project{Id: project_temp.Id, Name: project_temp.Name, Creator: &creator, Description: project_temp.Description, Star: project_temp.Star, Created: project_temp.Created, Updated: project_temp.Updated}
		favorites[i].Project = &project
	}
	return favorites
}

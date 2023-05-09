package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Outline struct {
	Id      int
	Outline string `orm:"type(text)"`
}

func GetOutline(id int) (Outline, error) {
	o := orm.NewOrm()
	outline := Outline{Id: id}
	err := o.Read(&outline)
	return outline, err
}

func CreateOutline(outline_str string) (Outline, error) {
	o := orm.NewOrm()
	outline := Outline{Outline: outline_str}
	_, err := o.Insert(&outline)
	return outline, err
}

func UpdateOutline(id int, outline string) (Outline, error) {
	o := orm.NewOrm()
	_outline := Outline{Id: id}
	err := o.Read(&_outline)
	if err != nil {
		return Outline{}, err
	}
	_outline.Outline = outline
	_, err = o.Update(&_outline)
	return _outline, err
}

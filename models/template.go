package models

import "github.com/beego/beego/v2/client/orm"

type Template struct {
	Id   int    `orm:"auto;pk"`
	Name string `orm:"size(100);unique"`
	// 模板内容，任意长度
	Cover      string `orm:"type(text)"`
	Thank      string `orm:"type(text)"`
	Transition string `orm:"type(text)"`
	Catalog_3  string `orm:"type(text)"`
	Catalog_4  string `orm:"type(text)"`
	Catalog_5  string `orm:"type(text)"`
	Content_2  string `orm:"type(text)"`
	Content_3  string `orm:"type(text)"`
	Content_4  string `orm:"type(text)"`
}

func GetAllTemplates() []Template {
	o := orm.NewOrm()
	var templates []Template
	o.QueryTable("template").All(&templates)
	return templates
}

func GetTemplate(id int) (Template, error) {
	o := orm.NewOrm()
	var template Template
	err := o.QueryTable("template").Filter("id", id).One(&template)
	return template, err
}

func GetTemplateByName(name string) (Template, error) {
	o := orm.NewOrm()
	var template Template
	err := o.QueryTable("template").Filter("name", name).One(&template)
	return template, err
}

func CreateTemplate(template Template) error {
	o := orm.NewOrm()
	_, err := o.Insert(&template)
	return err
}

package models

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

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
	Content_1  string `orm:"type(text)"`
	Content_2  string `orm:"type(text)"`
	Content_3  string `orm:"type(text)"`
	Content_4  string `orm:"type(text)"`
}

type JsonObject map[string]interface{}

func GetObj(obj string) JsonObject {
	var ret map[string]interface{}
	json.Unmarshal([]byte(obj), &ret)
	return ret
}

func GetJsonTemplate(template Template) JsonObject {
	return map[string]interface{}{
		"id":         template.Id,
		"name":       template.Name,
		"cover":      GetObj(template.Cover),
		"thank":      GetObj(template.Thank),
		"transition": GetObj(template.Transition),
		"catalog_3":  GetObj(template.Catalog_3),
		"catalog_4":  GetObj(template.Catalog_4),
		"catalog_5":  GetObj(template.Catalog_5),
		"content_1":  GetObj(template.Content_1),
		"content_2":  GetObj(template.Content_2),
		"content_3":  GetObj(template.Content_3),
		"content_4":  GetObj(template.Content_4),
	}
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

func GetTemplateDir(template_id int) string {
	return "static/template/" + strconv.Itoa(template_id)
}

func GetTemplateImageUrl(template_id int) string {
	image_path := "static/template/" + strconv.Itoa(template_id) + "/cover.png"
	//查看是否存在该文件
	_, err := os.Stat(image_path)
	if err != nil {
		return "http://{{server_ip}}:8080/_static/project/default.png"
	}
	res := "http://{{server_ip}}:8080/_" + image_path

	return res
}

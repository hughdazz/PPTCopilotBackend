package models

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
)

func jsonObject2string(obj JsonObject) string {
	ret, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(ret)
}

// 初始化数据表
func init() {
	// 获取环境变量
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	fmt.Println("MYSQL_HOST: ", mysqlHost)
	fmt.Println("MYSQL_PORT: ", mysqlPort)
	orm.RegisterDataBase("default", "mysql", "root:admin@tcp("+mysqlHost+":"+mysqlPort+")/now_db?charset=utf8&loc=Local")

	// 注册定义的model
	orm.RegisterModel(new(Outline))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Project))
	orm.RegisterModel(new(File))
	orm.RegisterModel(new(Favorite))
	orm.RegisterModel(new(Template))

	// 如果表不存在则创建表
	orm.RunSyncdb("default", false, true)

	// 读取scripts下jianyuezi.json文件
	jianyuezi, err := os.Open("./scripts/jianyuezi.json")
	if err != nil {
		fmt.Println(err)
	}
	var jianyueziJson []JsonObject
	// 解析json文件
	json.NewDecoder(jianyuezi).Decode(&jianyueziJson)
	template1 := Template{
		Name:       "简约紫",
		Cover:      jsonObject2string(jianyueziJson[0]),
		Transition: jsonObject2string(jianyueziJson[1]),
		Catalog_3:  jsonObject2string(jianyueziJson[2]),
		Catalog_4:  jsonObject2string(jianyueziJson[3]),
		Catalog_5:  jsonObject2string(jianyueziJson[4]),
		Content_1:  jsonObject2string(jianyueziJson[5]),
		Content_2:  jsonObject2string(jianyueziJson[6]),
		Content_3:  jsonObject2string(jianyueziJson[7]),
		Content_4:  jsonObject2string(jianyueziJson[8]),
		Thank:      jsonObject2string(jianyueziJson[9]),
	}
	CreateTemplate(template1)

	// 读取scripts下taikongren.json
	taikongren, err := os.Open("./scripts/taikongren.json")
	if err != nil {
		fmt.Println(err)
	}
	var taikongrenJson []JsonObject
	// 解析json文件
	json.NewDecoder(taikongren).Decode(&taikongrenJson)
	template2 := Template{
		Name:       "太空人",
		Cover:      jsonObject2string(taikongrenJson[0]),
		Transition: jsonObject2string(taikongrenJson[1]),
		Catalog_3:  jsonObject2string(taikongrenJson[2]),
		Catalog_4:  jsonObject2string(taikongrenJson[3]),
		Catalog_5:  jsonObject2string(taikongrenJson[4]),
		Content_1:  jsonObject2string(taikongrenJson[5]),
		Content_2:  jsonObject2string(taikongrenJson[6]),
		Content_3:  jsonObject2string(taikongrenJson[7]),
		Content_4:  jsonObject2string(taikongrenJson[8]),
		Thank:      jsonObject2string(taikongrenJson[9]),
	}
	CreateTemplate(template2)

	// 读取scripts下dangjian.json
	dangjian, err := os.Open("./scripts/dangjian.json")
	if err != nil {
		fmt.Println(err)
	}

	var dangjianJson []JsonObject
	// 解析json文件

	json.NewDecoder(dangjian).Decode(&dangjianJson)

	template3 := Template{
		Name:       "党建",
		Cover:      jsonObject2string(dangjianJson[0]),
		Transition: jsonObject2string(dangjianJson[1]),
		Catalog_3:  jsonObject2string(dangjianJson[2]),
		Catalog_4:  jsonObject2string(dangjianJson[3]),
		Catalog_5:  jsonObject2string(dangjianJson[4]),
		Content_1:  jsonObject2string(dangjianJson[5]),
		Content_2:  jsonObject2string(dangjianJson[6]),
		Content_3:  jsonObject2string(dangjianJson[7]),
		Content_4:  jsonObject2string(dangjianJson[8]),
		Thank:      jsonObject2string(dangjianJson[9]),
	}
	CreateTemplate(template3)

}

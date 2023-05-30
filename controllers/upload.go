package controllers

import (
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type UploadRouter struct {
	beego.Controller
}

func (c *UploadRouter) Post() {
	// 确定文件保存路径
	save_dir := c.GetString("savedir")

	err := os.MkdirAll(save_dir, 0777)
	if err != nil {
		c.Data["json"] = MakeResponse(Err, err.Error(), nil)
		c.ServeJSON()
		return
	}
	//如果save_dir以/结尾，则去掉/
	if save_dir[len(save_dir)-1] == '/' {
		save_dir = save_dir[:len(save_dir)-1]
	}

	// 将form-data中的文件保存到本地
	file_path := save_dir + "/" + c.GetString("savename")
	err = c.SaveToFile("uploadname", file_path)
	if err != nil {
		c.Data["json"] = MakeResponse(Err, err.Error(), nil)
		c.ServeJSON()
		return
	}

	c.Data["json"] = MakeResponse(OK, "success", nil)
	c.ServeJSON()
}

package project

import (
	"backend/controllers"
	"backend/models"
	"os"
	"strconv"
)

func (this *Controller) CreateFile() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	save_dir := models.GetSaveDir(id)
	// 递归创建目录
	err = os.MkdirAll(save_dir, 0777)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	save_name := this.GetString("savename")
	//如果文件以.json结尾，那么返回错误
	if len(save_name) > 5 && save_name[len(save_name)-5:] == ".json" {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "上传文件名不能以.json结尾", nil)
		this.ServeJSON()
		return
	}

	file_path := save_dir + "/" + save_name

	// 保存到数据库
	file, err := models.CreateFile(save_name, id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// 将form-data中的文件保存到本地
	err = this.SaveToFile("uploadname", file_path)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", file)
	this.ServeJSON()
}

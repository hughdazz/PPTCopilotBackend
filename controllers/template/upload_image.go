package template

import (
	"backend/controllers"
	"backend/models"
	"os"
	"strconv"
)

func (this *Controller) UploadImage() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	_, err = models.GetTemplate(id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "模板不存在", nil)
		this.ServeJSON()
		return
	}

	save_dir := models.GetTemplateDir(id)
	// 递归创建目录
	err = os.MkdirAll(save_dir, 0777)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	file_path := save_dir + "/" + "cover.png"

	// 将form-data中的文件保存到本地
	err = this.SaveToFile("uploadname", file_path)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}

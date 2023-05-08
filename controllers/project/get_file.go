package project

import (
	"backend/controllers"
	"backend/models"
	"encoding/base64"
	"io/ioutil"
	"strconv"
)

func (this *Controller) GetFile_() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}
	file_name := this.Ctx.Input.Param(":file_name")

	file, err := models.GetFile(file_name, id)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// 读取文件
	path := models.GetFilePathByName(file.Name, id)
	file_content, err := ioutil.ReadFile(path)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	encoded_content := base64.StdEncoding.EncodeToString(file_content)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", encoded_content)
	this.ServeJSON()
}

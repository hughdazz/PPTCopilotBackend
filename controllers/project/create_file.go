package project

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"strconv"
)

type CreateFileRequest struct {
	FileName *string `json:"file_name"`
}

func (this *Controller) CreateFile() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	var request CreateFileRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)
	if request.FileName == nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	file, err := models.CreateFile(*request.FileName, id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	save_dir := models.GetSaveDir(id)
	file_path := save_dir + "/" + file.Name

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

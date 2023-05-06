package project

import (
	"backend/controllers"
	"backend/models"
	"encoding/base64"
	"os"
	"strconv"
)

type UpdateFileRequest struct {
	FileContent *string `json:"file_content"`
}

func (this *Controller) UpdateFile() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	file_name := this.Ctx.Input.Param(":file_name")

	var request UpdateFileRequest
	this.ParseForm(&request)

	if request.FileContent == nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	decode_content, err := base64.StdEncoding.DecodeString(*request.FileContent)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	file_path := models.GetFilePathByName(file_name, id)

	os.WriteFile(file_path, decode_content, 0644)

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}

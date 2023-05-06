package project

import (
	"backend/controllers"
	"backend/models"
	"encoding/base64"
	"encoding/json"
	"os"
	"strconv"
)

type CreateFileRequest struct {
	FileName    *string `json:"file_name"`
	FileContent *string `json:"file_content"`
}

func (this *Controller) CreateFile() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	var request CreateFileRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)
	if request.FileName == nil || request.FileContent == nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	file, err := models.CreateFile(*request.FileName, id)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// base64 解码
	file_content, err := base64.StdEncoding.DecodeString(*request.FileContent)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	save_dir := models.GetSaveDir(id)
	file_path := save_dir + "/" + file.Name

	// 写入文件
	// 递归的创建目录
	err = os.MkdirAll(save_dir, 0755)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// 0644 如果文件不存在则创建，否则不创建
	err = os.WriteFile(file_path, file_content, 0644)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", file)
	this.ServeJSON()
}

package controllers

import (
	"backend/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type FileController struct {
	beego.Controller
}

type FileResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
}

func (this *FileController) Delete() {
	var delete_response FileResponse

	projectIdStr := this.Ctx.Input.Param(":project_id")
	projectId, _ := strconv.Atoi(projectIdStr)
	filename := this.Ctx.Input.Param(":file_name")

	res := models.DeleteFile(filename, projectId)
	delete_response.Id = projectId
	delete_response.Name = filename
	delete_response.Code = res
	if res == 1 {
		this.Ctx.Output.SetStatus(404)
		delete_response.Message = "文件不存在"
	} else if res == 2 {
		this.Ctx.Output.SetStatus(500)
		delete_response.Message = "文件结构删除错误"
	} else if res == 3 {
		this.Ctx.Output.SetStatus(500)
		delete_response.Message = "文件删除错误"
	} else {
		this.Ctx.Output.SetStatus(500)
		delete_response.Message = "文件删除正确"
	}

	this.Data["json"] = delete_response
	this.ServeJSON()
	return
}

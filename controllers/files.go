package controllers

import (
	"backend/models"
	"os"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type FilesController struct {
	beego.Controller
}

type FilesResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (this *FilesController) Post() {
	var file_response FilesResponse
	//获取项目id
	projectIdStr := this.Ctx.Input.Param(":project_id")
	projectId, _ := strconv.Atoi(projectIdStr)
	//获取文件
	file, header, err := this.GetFile("uploadname")
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		file_response.Code = 1
		file_response.Message = "文件上传错误，请重新添加！"
		this.Data["json"] = file_response
		this.ServeJSON()
		return
	}
	defer file.Close()
	//检查文件大小
	fileSize := header.Size
	if fileSize > 2*1024*1024 {
		this.Ctx.Output.SetStatus(500)
		file_response.Code = 1
		file_response.Message = "文件不能超过2MB大小，请重新添加！"
		this.Data["json"] = file_response
		this.ServeJSON()
		return
	}
	//查找该项目下有无重名文件
	_, res := models.GetFileInProjectByName(header.Filename, projectId)
	if res == 0 {
		this.Ctx.Output.SetStatus(500)
		file_response.Code = 1
		file_response.Message = "该项目下已经具有该文件，请重新添加或选择更新文件！"
		this.Data["json"] = file_response
		this.ServeJSON()
		return
	}

	// 创建保存文件的目录
	saveDir := "static/files/" + projectIdStr
	err = os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		file_response.Code = 1
		file_response.Message = "路径错误，请重新添加！"
		this.Data["json"] = file_response
		this.ServeJSON()
		return
	}

	filePath := saveDir + "/" + header.Filename

	// 保存文件
	err = this.SaveToFile("uploadname", filePath)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		file_response.Code = 1
		file_response.Message = "保存文件错误，请重新添加！"
		this.Data["json"] = file_response
		this.ServeJSON()
		return
	}
	_, err = models.NewFile(header.Filename, projectId)

	if err != nil {
		this.Ctx.Output.SetStatus(500)
		file_response.Code = 1
		file_response.Message = "添加文件错误，请重新添加！"
		this.Data["json"] = file_response
		this.ServeJSON()
		return
	}
	err = models.UpdataProject(projectId)
	this.Ctx.Output.SetStatus(200)
	file_response.Code = 0
	file_response.Message = "上传文件成功"
	this.Data["json"] = file_response
	this.ServeJSON()
	return

}

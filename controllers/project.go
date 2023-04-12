package controllers

import (
	"backend/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type ProjectController struct {
	beego.Controller
}

type DeleteProjectResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Id      int    `json:"id"`
}

func (this *ProjectController) Delete() {
	var delete_response DeleteProjectResponse

	projectIdStr := this.Ctx.Input.Param(":project_id")
	projectId, _ := strconv.Atoi(projectIdStr)

	err := models.DeleteProject(projectId)

	if err != nil {
		delete_response.Code = 1
		delete_response.Message = "删除失败"
		delete_response.Id = projectId

	} else {
		delete_response.Code = 0
		delete_response.Message = "删除成功"
		delete_response.Id = projectId
	}

	this.Data["json"] = delete_response
	this.ServeJSON()
}

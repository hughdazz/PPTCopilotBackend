package project

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

type UpdateProjectRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (this *Controller) UpdateProject() {
	var request UpdateProjectRequest
	this.ParseForm(&request)

	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}
	if request.Name != nil {
		_, err := models.UpdateProjectName(id, *request.Name)
		if err != nil {
			this.Ctx.Output.SetStatus(500)
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
	}
	if request.Description != nil {
		_, err := models.UpdateProjectDescription(id, *request.Description)
		if err != nil {
			this.Ctx.Output.SetStatus(500)
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}

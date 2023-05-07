package project

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetAll() {
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", models.GetAllProjects())
	this.ServeJSON()
}
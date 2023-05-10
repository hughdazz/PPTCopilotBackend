package project

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetAll() {

	projects := models.GetAllProjects()

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", models.RefactProjects(projects))
	this.ServeJSON()
}

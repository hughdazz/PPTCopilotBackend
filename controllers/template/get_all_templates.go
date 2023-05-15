package template

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetAllTemplates() {
	templates := models.GetAllTemplates()

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "ok", templates)
	this.ServeJSON()
}

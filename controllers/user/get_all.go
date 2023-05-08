package user

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetAll() {

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", models.GetAllUsers())
	this.ServeJSON()
}

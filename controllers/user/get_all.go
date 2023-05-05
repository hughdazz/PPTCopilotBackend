package user

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetAll() {
	this.Data["json"] = controllers.MakeResponse(200, "success", models.GetAllUsers())
	this.ServeJSON()
}

package user

import (
	"backend/controllers"
	"backend/models"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (this *Controller) CreateUser() {
	var request CreateUserRequest
	this.ParseForm(&request)
	user, err := models.CreateUser(request.Username, request.Password, request.Email)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", user)
	this.ServeJSON()
}

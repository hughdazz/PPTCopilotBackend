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
		this.Data["json"] = controllers.MakeResponse(401, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(200, "success", user)
	this.ServeJSON()
}

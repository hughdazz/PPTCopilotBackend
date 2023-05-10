package user

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
)

type CreateUserRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (this *Controller) CreateUser() {
	var request CreateUserRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)
	if request.Username == nil || request.Email == nil || request.Password == nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	user, err := models.CreateUser(*request.Username, *request.Password, *request.Email)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", user)
	this.ServeJSON()
}

package user

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

type UpdateUserRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (this *Controller) UpdateUser() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	var request UpdateUserRequest
	this.ParseForm(&request)
	if request.Username != nil {
		err = models.UpdateUserUsername(id, *request.Username)
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
	}
	if request.Email != nil {
		err = models.UpdateUserEmail(id, *request.Email)
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
	}
	if request.Password != nil {
		err = models.UpdateUserPassword(id, *request.Password)
		if err != nil {
			this.Ctx.Output.SetStatus(401)
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}

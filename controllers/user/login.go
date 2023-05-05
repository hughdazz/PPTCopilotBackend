package user

import (
	"backend/controllers"
	"backend/models"
)

type LoginRequest struct {
	Username_or_email string `json:"username_or_email"`
	Password          string `json:"password"`
}

func (this *Controller) Login() {
	var login_request LoginRequest
	this.ParseForm(&login_request)
	user, err := models.VerifyUser(login_request.Username_or_email, login_request.Password)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	tokenString := models.CreateToken(user.Id)
	// 设置在响应头中
	this.Ctx.SetCookie("token", tokenString, "3600", "/")

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", user)
	this.ServeJSON()
}

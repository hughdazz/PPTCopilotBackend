package user

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
)

type LoginRequest struct {
	Username_or_email *string `json:"username_or_email"`
	Password          *string `json:"password"`
}

func (this *Controller) Login() {
	var request LoginRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	if request.Username_or_email == nil || request.Password == nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	user, err := models.VerifyUser(*request.Username_or_email, *request.Password)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	tokenString := models.CreateToken(user.Id)
	// 设置在响应头中
	this.Ctx.SetCookie("token", tokenString, "3600", "/")

	this.Data["json"] = map[string]interface{}{
		"code":    controllers.OK,
		"message": "success",
		"token":   tokenString, "data": user}
	this.ServeJSON()
}

package controllers

import (
	"backend/models"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	beego.Controller
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Repassword string `json:"repassword"`
}

type RegisterResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (this *RegisterController) Post() {
	// 打印请求json
	// 解析json数据
	var Register_request RegisterRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&Register_request)
	username := Register_request.Username
	email := Register_request.Email
	password := Register_request.Password
	// repassword := Register_request.Repassword
	// 添加新用户
	_, err := models.CreateUser(username, password, email)
	var Register_response RegisterResponse

	if err != nil {
		// 验证失败，返回错误信息
		this.Ctx.Output.SetStatus(401)
		Register_response.Code = 1
		Register_response.Message = "注册错误"
		this.Data["json"] = Register_response
		this.ServeJSON()
		return
	}

	this.Ctx.Output.SetStatus(200)
	Register_response.Code = 0
	Register_response.Message = "注册成功"
	this.Data["json"] = Register_response
	this.ServeJSON()
	return

}

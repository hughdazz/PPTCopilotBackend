package controllers

import (
	"backend/models"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

type LoginRequest struct {
	Username_or_email string `json:"username_or_email"`
	Password          string `json:"password"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (this *LoginController) Post() {
	// 打印请求json

	// 解析json数据
	var login_request LoginRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&login_request)
	username_or_email := login_request.Username_or_email
	password := login_request.Password

	// 验证用户名和密码是否正确
	user, err := models.VerifyUser(username_or_email, password)

	var login_response LoginResponse

	if err != nil {
		// 验证失败，返回错误信息
		this.Ctx.Output.SetStatus(401)
		login_response.Code = 1
		login_response.Message = "用户名或密码错误"
		this.Data["json"] = login_response
		this.ServeJSON()
		return
	}
	// 成功，设置session
	// this.SetSession("user_id", user.Id)

	// 创建token
	tokenString := models.CreateToken(user.Id)
	this.Ctx.Output.Header("token", tokenString)
	this.Ctx.SetCookie("token", tokenString, "3600", "/")

	this.Ctx.Output.SetStatus(200)
	login_response.Code = 0
	login_response.Message = "登录成功"
	login_response.Token = tokenString
	this.Data["json"] = login_response

	this.ServeJSON()
}

// 登录模块

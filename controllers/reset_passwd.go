package controllers

import (
	"backend/models"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type ResetPasswordController struct {
	beego.Controller
}

type ResetPasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

func (this *ResetPasswordController) Post() {
	// 获取请求中的json数据
	var request ResetPasswordRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	user_id := this.GetSession("user_id")

	if user_id == nil {
		// 用户未登录
		this.Redirect("/login", 302)
		return
	}

	user, err := models.GetUser(user_id.(int))
	email := user.Email
	password := request.Password
	newPassword := request.NewPassword

	// 获取email失败，返回错误信息
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "获取邮箱失败"}
		this.ServeJSON()
		return
	}

	// 验证用户信息
	_, err = models.VerifyUser(email, password)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名或密码错误"}
		this.ServeJSON()
		return
	}

	// 更新密码
	err = models.UpdatePassword(email, newPassword)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = map[string]interface{}{"code": 2, "message": "修改密码失败"}
		this.ServeJSON()
		return
	}

	// 返回成功信息
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = map[string]interface{}{"code": 0, "message": "修改密码成功"}
	this.ServeJSON()
}

package controllers

import (
	"backend/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type CheckLoginController struct {
	beego.Controller
}

type CheckLoginResponse struct {
	IsAuthorized int    `json:"isAuthorized"`
	Username     string `json:"username"`
	Email        string `json:"email"`
}

func (this *CheckLoginController) Get() {
	var checklogin CheckLoginResponse
	user_id := this.GetSession("user_id")
	fmt.Printf("check_login\n")
	if user_id == nil {
		// 用户未登录
		this.Ctx.Output.SetStatus(401)
		checklogin.IsAuthorized = 0
		this.Data["json"] = checklogin
		this.ServeJSON()
		return
	}
	fmt.Printf("has_login\n")
	user, err := models.GetUser(user_id.(int))
	// 获取失败，返回错误信息
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		checklogin.IsAuthorized = 0
		this.Data["json"] = checklogin
		this.ServeJSON()
		return
	}
	this.Ctx.Output.SetStatus(200)
	checklogin.IsAuthorized = 1
	checklogin.Username = user.Username
	checklogin.Email = user.Email
	this.Data["json"] = checklogin
	this.ServeJSON()
}

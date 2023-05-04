package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LogoutController struct {
	beego.Controller
}

type LogoutResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (this *LogoutController) Delete() {
	// this.DelSession("user_id")

	// 会有登出失败的情况吗？

	this.Ctx.Output.SetStatus(200)
	var logout_response LogoutResponse
	logout_response.Code = 0
	logout_response.Message = "登出成功"
	this.Data["json"] = logout_response
	this.ServeJSON()
}

package user

import "backend/controllers"

func (this *Controller) Logout() {
	// 使浏览器删除 cookie
	this.Ctx.SetCookie("token", "", "-1", "/")

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}

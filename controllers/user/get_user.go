package user

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GetUser() {
	// 获取路由参数
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	user, err := models.GetUser(id)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", user)
	this.ServeJSON()
}

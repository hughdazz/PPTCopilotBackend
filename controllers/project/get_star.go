package project

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GetStar() {
	// 获取cookie
	token, err := this.Ctx.Request.Cookie("token")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	id := models.GetUserId(token.Value)

	id_ := this.Ctx.Input.Param(":id")
	project_id, err := strconv.Atoi(id_)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	_, res, err := models.GetFavorite(id, project_id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", res)
	this.ServeJSON()

}

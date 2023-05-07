package project

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) DeleteProject() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	err = models.DeleteProject(id)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()
}

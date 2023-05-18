package project

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GetProject() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	project, err := models.GetProject(id)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	project = models.RefactProject(project)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", project)
	this.ServeJSON()
}

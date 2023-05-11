package project

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GetFiles() {
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	files, err := models.GetFiles(id)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	files = models.RefactFiles(files)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", files)
	this.ServeJSON()
}

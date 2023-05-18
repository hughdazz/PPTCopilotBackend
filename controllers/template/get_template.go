package template

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GetTemplate() {
	// :id
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	template, err := models.GetTemplate(id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	Jsontemplate := models.GetJsonTemplate(template)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "ok", Jsontemplate)
	this.ServeJSON()
}

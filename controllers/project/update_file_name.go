package project

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"strconv"
)

type UpdateFileRequest struct {
	OldName *string `json:"old_name"`
	Name    *string `json:"new_name"`
}

func (this *Controller) UpdateFileName() {
	var request UpdateFileRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	if request.OldName == nil || request.Name == nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	file, err := models.UpdateFileName(id, *request.OldName, *request.Name)
	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", file)
	this.ServeJSON()

}

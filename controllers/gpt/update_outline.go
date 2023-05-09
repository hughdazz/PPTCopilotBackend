package gpt

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"strconv"
)

type UpdateOutlineRequest struct {
	Outline string `json:"outline"`
}

func (this *Controller) UpdateOutline() {
	var request UpdateOutlineRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)
	id_ := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}
	_, err = models.UpdateOutline(id, request.Outline)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", nil)
	this.ServeJSON()

}

package project

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
)

type CreateProjectRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (this *Controller) CreateProject() {
	var request CreateProjectRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	if request.Name == nil || request.Description == nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}
	// 获取cookie
	token, err := this.Ctx.Request.Cookie("token")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	id := models.GetUserId(token.Value)

	//判断项目名是否为空串、或仅由空格和制表符组成的串
	if len(*request.Name) == 0 {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "项目名不能为空", nil)
		this.ServeJSON()
		return
	}
	var is_space bool = true
	for _, v := range *request.Name {
		if v != ' ' && v != '\t' {
			is_space = false
			break
		}
	}
	if is_space {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "项目名不能为空", nil)
		this.ServeJSON()
		return
	}

	project, err := models.CreateProject(*request.Name, *request.Description, id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", project)
	this.ServeJSON()
}

package controllers

import (
	"backend/models"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type ProjectController struct {
	beego.Controller
}

type NewProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewProjectResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetProjectsResponse struct {
	Code     int              `json:"code"`
	Projects []models.Project `json:"projects"`
	Message  string           `json:"message"`
}

// Post方法
// 新建项目
func (this *ProjectController) Post() {
	var new_project_response NewProjectResponse
	// 获取user_id
	user_id := this.GetSession("user_id")
	if user_id == nil {
		this.Ctx.Output.SetStatus(401)
		new_project_response.Code = 1
		new_project_response.Message = "请先登录"
		this.Data["json"] = new_project_response
		this.ServeJSON()
		return
	}
	var new_project_request NewProjectRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&new_project_request)
	name := new_project_request.Name
	description := new_project_request.Description

	// 创建项目
	_, err := models.NewProject(name, description, user_id.(int))

	if err != nil {
		this.Ctx.Output.SetStatus(500)
		new_project_response.Code = 1
		new_project_response.Message = "创建项目失败"
		this.Data["json"] = new_project_response
		this.ServeJSON()
		return
	}

	this.Ctx.Output.SetStatus(200)
	new_project_response.Code = 0
	new_project_response.Message = "创建项目成功"
	this.Data["json"] = new_project_response
	this.ServeJSON()
}

func (this *ProjectController) Get() {
	var get_projects_response GetProjectsResponse
	// 获取user_id
	user_id := this.GetSession("user_id")
	if user_id == nil {
		this.Ctx.Output.SetStatus(401)
		get_projects_response.Code = 1
		get_projects_response.Message = "请先登录"
		this.Data["json"] = get_projects_response
		this.ServeJSON()
		return
	}

	// 获取项目列表
	projects, err := models.GetProjects(user_id.(int))
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		get_projects_response.Code = 1
		get_projects_response.Message = "获取项目列表失败"
		this.Data["json"] = get_projects_response
		this.ServeJSON()
		return
	}

	this.Ctx.Output.SetStatus(200)
	get_projects_response.Code = 0
	get_projects_response.Message = "获取项目列表成功"
	get_projects_response.Projects = projects
	this.Data["json"] = get_projects_response
	this.ServeJSON()
}

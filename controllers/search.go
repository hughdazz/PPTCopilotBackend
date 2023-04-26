package controllers

import (
	"backend/models"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type SearchController struct {
	beego.Controller
}

type SearchProjectsResponse struct {
	Code     int              `json:"code"`
	Search   []string         `json:"search"`
	Projects []models.Project `json:"projects"`
	Message  string           `json:"message"`
}

func (this *SearchController) Get() {
	var Search_projects_response SearchProjectsResponse

	filterWords := this.GetString("filter_words")

	// 拆分关键词
	keywords := strings.Split(filterWords, " ")

	// 查询项目
	projects, err := models.SearchProjects(keywords)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		Search_projects_response.Code = 1
		Search_projects_response.Message = "获取项目列表失败"

	} else {
		this.Ctx.Output.SetStatus(200)
		Search_projects_response.Code = 0
		Search_projects_response.Message = "获取项目列表成功"
		Search_projects_response.Projects = projects
	}
	Search_projects_response.Search = keywords
	this.Data["json"] = Search_projects_response
	this.ServeJSON()
	return
}

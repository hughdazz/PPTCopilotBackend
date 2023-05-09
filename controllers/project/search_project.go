package project

import (
	"backend/controllers"
	"backend/models"
	"strings"
)

func (this *Controller) SearchProject() {
	filterWords := this.GetString("filter_words")

	// 拆分关键词
	keywords := strings.Split(filterWords, " ")

	// 查询项目
	projects_temp, err := models.SearchProjects(keywords)

	// 重新组装项目
	projects := models.RefactProjects(projects_temp)

	if err != nil {

		this.Data["json"] = controllers.MakeResponse(controllers.Err, "获取项目列表失败", nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "获取项目列表成功", projects)
	this.ServeJSON()

}

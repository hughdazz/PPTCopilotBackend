package template

import (
	"backend/controllers"
	"backend/models"
)

type InfoResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ImageUrl string `json:"imageUrl"`
}

func (this *Controller) GetTemplatesInfo() {
	templates := models.GetAllTemplates()

	// 转换为json
	jsonTemplates := make([]InfoResponse, len(templates))
	for i, template := range templates {
		jsonTemplates[i] = InfoResponse{
			Id:       template.Id,
			Title:    template.Name,
			ImageUrl: models.GetTemplateImageUrl(template.Id),
		}
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "ok", jsonTemplates)
	this.ServeJSON()
}

package template

import (
	"backend/controllers"
	"backend/models"
)

type InfoResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

func (this *Controller) GetTemplatesInfo() {
	templates := models.GetAllTemplates()

	// 转换为json
	jsonTemplates := make([]InfoResponse, len(templates))
	for i, template := range templates {
		jsonTemplates[i] = InfoResponse{
			Id:       template.Id,
			Name:     template.Name,
			ImageUrl: models.GetTemplateImageUrl(template.Id),
		}
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "ok", jsonTemplates)
	this.ServeJSON()
}

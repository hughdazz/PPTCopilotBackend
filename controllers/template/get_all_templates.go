package template

import (
	"backend/controllers"
	"backend/models"
)

func (this *Controller) GetAllTemplates() {
	templates := models.GetAllTemplates()

	// 转换为json
	jsonTemplates := make([]map[string]interface{}, len(templates))
	for i, template := range templates {
		jsonTemplates[i] = models.GetJsonTemplate(template)
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "ok", jsonTemplates)
	this.ServeJSON()
}

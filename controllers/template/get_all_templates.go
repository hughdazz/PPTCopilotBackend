package template

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
)

func getObj(obj string) JsonObject {
	var ret map[string]interface{}
	json.Unmarshal([]byte(obj), &ret)
	return ret
}

func GetJsonTemplate(template models.Template) JsonObject {
	return map[string]interface{}{
		"id":         template.Id,
		"name":       template.Name,
		"cover":      getObj(template.Cover),
		"thank":      getObj(template.Thank),
		"transition": getObj(template.Transition),
		"catalog_3":  getObj(template.Catalog_3),
		"catalog_4":  getObj(template.Catalog_4),
		"catalog_5":  getObj(template.Catalog_5),
		"content_2":  getObj(template.Content_2),
		"content_3":  getObj(template.Content_3),
		"content_4":  getObj(template.Content_4),
	}
}

func (this *Controller) GetAllTemplates() {
	templates := models.GetAllTemplates()

	// 转换为json
	jsonTemplates := make([]map[string]interface{}, len(templates))
	for i, template := range templates {
		jsonTemplates[i] = GetJsonTemplate(template)
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "ok", jsonTemplates)
	this.ServeJSON()
}

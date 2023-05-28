package template

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
)

type CreateRequest struct {
	Name     string              `json:"name"`
	Template []models.JsonObject `json:"template"`
	// Cover      models.JsonObject `json:"cover"`
	// Thank      models.JsonObject `json:"thank"`
	// Transition models.JsonObject `json:"transition"`
	// Catalog_3  models.JsonObject `json:"catalog-3"`
	// Catalog_4  models.JsonObject `json:"catalog-4"`
	// Catalog_5  models.JsonObject `json:"catalog-5"`
	// Content_1  models.JsonObject `json:"content-1"`
	// Content_2  models.JsonObject `json:"content-2"`
	// Content_3  models.JsonObject `json:"content-3"`
	// Content_4  models.JsonObject `json:"content-4"`
}

type PPTPages struct {
	Pages []models.JsonObject `json:"pages"`
}

func getRet(this *Controller, obj models.JsonObject) string {
	ret, err := json.Marshal(obj)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return ""
	}
	return string(ret)
}
func (this *Controller) CreateTemplate() {
	var req CreateRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&req)

	template := models.Template{
		Name:       req.Name,
		Cover:      getRet(this, req.Template[0]),
		Transition: getRet(this, req.Template[1]),
		Catalog_3:  getRet(this, req.Template[2]),
		Catalog_4:  getRet(this, req.Template[3]),
		Catalog_5:  getRet(this, req.Template[4]),
		Content_1:  getRet(this, req.Template[5]),
		Content_2:  getRet(this, req.Template[6]),
		Content_3:  getRet(this, req.Template[7]),
		Content_4:  getRet(this, req.Template[8]),
		Thank:      getRet(this, req.Template[9]),
	}
	err := models.CreateTemplate(template)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "ok", nil)
	this.ServeJSON()
}

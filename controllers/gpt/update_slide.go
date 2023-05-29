package gpt

import (
	"backend/conf"
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"strings"
)

type UpdateSlideRequest struct {
	Prompt string `json:"prompt"`
	Slide  string `json:"slide"`
}

func (this *Controller) UpdateSlide() {
	var request UpdateSlideRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	template := conf.GetUpdateSinglePromptTemplate()
	template = strings.ReplaceAll(template, "{{prompt}}", request.Prompt)
	template = strings.ReplaceAll(template, "{{slide}}", request.Slide)

	updated_slide, err := RequestGpt(template, SectionXML{}) // <section></section>
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	updated_slide = models.RefactXML(updated_slide)

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", updated_slide)
	this.ServeJSON()
}

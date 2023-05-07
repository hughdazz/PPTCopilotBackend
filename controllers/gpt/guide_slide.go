package gpt

import (
	"backend/conf"
	"backend/controllers"
	"encoding/json"
	"strings"
)

type GuideSlideRequest struct {
	Outline string `json:"outline"`
}

func (this *Controller) GuideSlide() {
	var request GuideSlideRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	template := conf.GetGuideSinglePromptTemplate()
	template = strings.ReplaceAll(template, "{{outline}}", request.Outline)

	guide_slide, err := RequestGpt(template)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", guide_slide)
	this.ServeJSON()

}

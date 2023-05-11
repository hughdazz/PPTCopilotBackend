package gpt

import (
	"backend/conf"
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"strings"
)

type GenOutlineRequest struct {
	Topic   string `json:"topic"`
	Sponsor string `json:"sponsor"`
}

func (this *Controller) GenOutline() {
	var request GenOutlineRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	prompt := conf.GetOutlinePromptTemplate()
	prompt = strings.ReplaceAll(prompt, "{{topic}}", request.Topic)
	prompt = strings.ReplaceAll(prompt, "{{sponsor}}", request.Sponsor)

	outline_str, err := RequestGpt(prompt, SlidesXML{}) //<slide></slide>
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), outline_str)
		this.ServeJSON()
		return
	}

	// 替换换行符
	outline_str = strings.ReplaceAll(outline_str, "\n", "")
	outline_str = strings.ReplaceAll(outline_str, "\\\n", "")

	outline, err := models.CreateOutline(outline_str)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), outline)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", outline)
	this.ServeJSON()
}

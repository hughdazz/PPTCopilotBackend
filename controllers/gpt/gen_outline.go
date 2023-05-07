package gpt

import (
	"backend/conf"
	"backend/controllers"
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

	outline, err := RequestGpt(prompt)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", outline)
	this.ServeJSON()
}

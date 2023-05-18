package gpt

import (
	"backend/controllers"
	"backend/models"
	"strconv"
)

func (this *Controller) GenPPT() {

	//从form-data获取outline_id和template_id
	outline_id_ := this.GetString("outline_id")
	outline_id, err := strconv.Atoi(outline_id_)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}
	template_id_ := this.GetString("template_id")
	template_id, err := strconv.Atoi(template_id_)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	//从数据库获取outline和template
	outline, err := models.GetOutline(outline_id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	template, err := models.GetTemplate(template_id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), template)
		this.ServeJSON()
		return
	}

	// 获取所有的ContentSections
	content_sections, err := models.GetContentSections(outline.Outline)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// 所有的ContentSection进行guide_slide
	guide_slides := make([]string, 0)
	for _, content_section := range content_sections {
		guide_slide, err := GuideContentSection(content_section)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
			this.ServeJSON()
			return
		}
		guide_slides = append(guide_slides, guide_slide)
	}

	// 将outline.Outline中的所有的ContentSection替换为guide_slide
	resultxml, err := models.RefactContentSections(outline.Outline, guide_slides)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	var res []string

	res, err = models.GenPPT(resultxml, template)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
		this.ServeJSON()
		return
	}

	res, err = models.GenPPT(resultxml, template)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
		this.ServeJSON()
		return
	}

	JsonRes := make([]models.JsonObject, len(res))
	for i, _ := range res {
		JsonRes[i] = models.GetObj(res[i])
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", JsonRes)
	this.ServeJSON()

}

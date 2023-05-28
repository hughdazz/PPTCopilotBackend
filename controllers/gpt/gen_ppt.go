package gpt

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"fmt"
)

type GenPPTRequest struct {
	OutlineId  int    `json:"outline_id"`
	TemplateId int    `json:"template_id"`
	ProjectId  int    `json:"project_id"`
	FileName   string `json:"file_name"`
}

func (this *Controller) GenPPT() {
	var request GenPPTRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	outlineId := request.OutlineId
	templateId := request.TemplateId
	projectId := request.ProjectId
	fileName := request.FileName

	//从数据库获取outline和template
	outline, err := models.GetOutline(outlineId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	template, err := models.GetTemplate(templateId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), template)
		this.ServeJSON()
		return
	}

	debug := 0
	if debug == 1 {
		resultxml := `<slides>
		<section class='cover'>
			<p>我为什么玩明日方舟</p>
			<p>汇报人：dhf</p>
		</section>
		<section
			class='catalog'>
			<p>目录</p>
			<p>1. 玩法简介</p>
			<p>2. 游戏特色</p>
			<p>3. 社交互动</p>
			<p>4. 个人体验</p>
			<p>5. 结束语</p>
		</section>
		<section class='content'>
			<p>社交互动</p>
			<p>1. 分享自己在游戏中的心得体会有助于与其他玩家建立更紧密的联系，增强游戏体验。</p>
			<p>2. 参与游戏社区的互动活动，不仅可以赢取奖励，还能结交志同道合的朋友。</p>
			<p>3. 玩家之间的互动是游戏中不可或缺的一部分，可以互相帮助、交流游戏心得、组队挑战副本等。</p>
			<p>4. 玩家之间的互动是游戏中不可或缺的一部分，可以互相帮助、交流游戏心得、组队挑战副本等。</p>
			<p>5. 玩家之间的互动是游戏中不可或缺的一部分，可以互相帮助、交流游戏心得、组队挑战副本等。</p>
		</section>
	</slides>`

		var res []string

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

		file, err := models.CreateFile(fileName, projectId)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), file)
			this.ServeJSON()
			return
		}
		err = models.SaveJsonsToFile(JsonRes, fileName, projectId)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
			this.ServeJSON()
			return
		}

		this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", JsonRes)
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
		fmt.Println(guide_slide)
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

	JsonRes := make([]models.JsonObject, len(res))
	for i, _ := range res {
		JsonRes[i] = models.GetObj(res[i])
	}

	file, err := models.CreateFile(fileName, projectId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), file)
		this.ServeJSON()
		return
	}
	err = models.SaveJsonsToFile(JsonRes, fileName, projectId)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", JsonRes)
	this.ServeJSON()

}

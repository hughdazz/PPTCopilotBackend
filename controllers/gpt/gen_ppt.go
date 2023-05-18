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
	project_id_ := this.GetString("project_id")
	project_id, err := strconv.Atoi(project_id_)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	file_name := this.GetString("file_name")
	file_name = file_name + ".json"

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

		file, err := models.CreateFile(file_name, project_id)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), file)
			this.ServeJSON()
			return
		}
		err = models.SaveJsonsToFile(JsonRes, file_name, project_id)
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

	file, err := models.CreateFile(file_name, project_id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), file)
		this.ServeJSON()
		return
	}
	err = models.SaveJsonsToFile(JsonRes, file_name, project_id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), res)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", JsonRes)
	this.ServeJSON()

}

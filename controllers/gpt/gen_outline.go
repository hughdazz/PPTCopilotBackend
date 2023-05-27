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

	outline_str := ``
	debug := 0
	if debug == 1 {
		outline_str = `<slides>
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
	} else {
		var err error
		outline_str, err = RequestGpt(prompt, SlidesXML{}) //<slide></slide>
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), outline_str)
			this.ServeJSON()
			return
		}
	}

	outline, err := models.CreateOutline(outline_str)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), outline)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", outline)
	this.ServeJSON()
}

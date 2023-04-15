package controllers

import (
	"backend/conf"
	"backend/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
	"sync"
)

type GptController struct {
	beego.Controller
}

var chatContextMap = make(map[string]models.ChatContext)

type Config struct {
	data map[string]string
	lock sync.Mutex
}

type GetCatalogRequest struct {
	Prompt string `json:"prompt"`
}

type GetCatalogResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    string `json:"data"`
}

// GetCatalog 获取生成的PPT大纲
func (c *GptController) GetCatalog() {
	var getCataLogRequest GetCatalogRequest
	//corePrompt := c.Ctx.Request.Form.Get("prompt")
	json.NewDecoder(c.Ctx.Request.Body).Decode(&getCataLogRequest)
	var getCatalogResponse GetCatalogResponse
	// get userid as string
	userId := c.GetSession("user_id")
	if userId == nil {
		c.Ctx.Output.SetStatus(401)
		getCatalogResponse.Code = 1
		getCatalogResponse.Message = "用户未登录"
		c.Data["json"] = getCatalogResponse
		c.ServeJSON()
		return
	}
	user, err := models.GetUser(userId.(int))
	if err != nil {
		c.Ctx.Output.SetStatus(401)
		getCatalogResponse.Code = 1
		getCatalogResponse.Message = "用户未登录"
		c.Data["json"] = getCatalogResponse
		c.ServeJSON()
		return
	}
	apikey, err := models.GetApiKeyByUserId(user.Id)
	if apikey == "" {
		c.Ctx.Output.SetStatus(401)
		getCatalogResponse.Code = 1
		getCatalogResponse.Message = "用户未设置apikey"
		c.Data["json"] = getCatalogResponse
		c.ServeJSON()
		return
	}

	//	prompt := `
	//⽤⼾: ⽣成介绍 {{topic}} 的⼤纲,汇报人叫做{{user_name}}， 仿照以下格式， 不要带上任何注释性⽂字
	//<slides>
	//
	//<section class='封面'>
	//    <!--第一张幻灯片-->
	//    <h1>《论语》</h1>
	//    <p>汇报人：</p>
	//</section>
	//<section class='目录页'>
	//    <!--只用给出目录项-->
	//    <h1> 目录 </h1>
	//    <h2> 目录项1 </h2>
	//    <h2> 目录项2 </h2>
	//    <h2> 目录项3 </h2>
	//</section>
	//<section class='内容'>
	//    <!--只用给出一个标题和一个内容概要即可-->
	//    <hl> 标题 </hl>
	//    <p>内容概要</p>
	//</section>
	//<!--更多的幻灯片・・・-->
	//</slides>
	//    `
	prompt := conf.GetCatalogPromptTemplate()
	prompt = strings.ReplaceAll(prompt, "{{topic}}", getCataLogRequest.Prompt)
	prompt = strings.ReplaceAll(prompt, "{{user_name}}", user.Username)

	chatContext, _ := chatContextMap[string(user.Id)]
	generatedText, err := chatContext.RequestGptService(prompt, apikey, 0.7)

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		getCatalogResponse.Code = 0
		getCatalogResponse.Message = err.Error()
		c.Data["json"] = getCatalogResponse
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	getCatalogResponse.Code = 0
	getCatalogResponse.Message = "success"
	getCatalogResponse.Data = generatedText
	c.Data["json"] = getCatalogResponse
	c.ServeJSON()
	return
}

type UpdateSingleRequest struct {
	Prompt string `json:"prompt"`
	PPTXml string `json:"ppt_xml"`
}

type UpdateSingleResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    string `json:"data"`
}

// UpdateSides 修改单页ppt
func (c *GptController) UpdateSides() {
	var updateSingleRequest UpdateSingleRequest
	var updateSingleResponse UpdateSingleResponse

	json.NewDecoder(c.Ctx.Request.Body).Decode(&updateSingleRequest)

	userId := c.GetSession("user_id")
	if userId == nil {
		c.Ctx.Output.SetStatus(401)
		updateSingleResponse.Code = 1
		updateSingleResponse.Message = "用户未登录"
		c.Data["json"] = updateSingleResponse
		c.ServeJSON()
		return
	}
	user, err := models.GetUser(userId.(int))
	if err != nil {
		c.Ctx.Output.SetStatus(401)
		updateSingleResponse.Code = 1
		updateSingleResponse.Message = "用户未登录"
		c.Data["json"] = updateSingleResponse
		c.ServeJSON()
		return
	}
	apikey, err := models.GetApiKeyByUserId(user.Id)
	if apikey == "" {
		c.Ctx.Output.SetStatus(401)
		updateSingleResponse.Code = 1
		updateSingleResponse.Message = "用户未设置apikey"
		c.Data["json"] = updateSingleResponse
		c.ServeJSON()
		return
	}

	//prompt := updateSingleRequest.Prompt
	//pptXml := updateSingleRequest.PPTXml

	//	prompt := `
	//你将接收用户的指令去处理一个ppt，ppt的表示方式是xml格式的字符串。
	//用户的指令是：{{prompt}}。请同时帮我把这个ppt中双括号{{}}括起来的部分替换为与{{prompt}}相关的内容；
	//要处理的ppt是：{{ppt_xml}}；
	//请你按照用户的指令对以xml格式表示的ppt进行处理。
	//注意：你只用返回处理后的xml格式字符串，且不要带任何注释，谢谢！
	//说明：用双括号{{}}括起来的文本表示模板信息，是你一定要替换的部分`
	prompt := conf.GetUpdateSinglePromptTemplate()

	prompt = strings.ReplaceAll(prompt, "{{prompt}}", updateSingleRequest.Prompt)
	prompt = strings.ReplaceAll(prompt, "{{ppt_xml}}", updateSingleRequest.PPTXml)

	chatContext, _ := chatContextMap[string(user.Id)]
	generatedText, err := chatContext.RequestGptService(prompt, apikey, 0.7)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		updateSingleResponse.Code = 1
		updateSingleResponse.Message = "gpt服务出错" + err.Error()
		c.Data["json"] = updateSingleResponse
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	updateSingleResponse.Code = 0
	updateSingleResponse.Message = "success"
	updateSingleResponse.Data = generatedText
	c.Data["json"] = updateSingleResponse
	c.ServeJSON()
	return

}

type SetApiKeyRequest struct {
	ApiKey string `json:"api_key"`
}

type SetApiKeyResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (c *GptController) SetApiKey() {
	var setApiKeyRequest SetApiKeyRequest
	var setApiKeyResponse SetApiKeyResponse
	json.NewDecoder(c.Ctx.Request.Body).Decode(&setApiKeyRequest)
	apiKey := setApiKeyRequest.ApiKey
	if apiKey == "" {
		c.Ctx.Output.SetStatus(400)
		setApiKeyResponse.Code = 1
		setApiKeyResponse.Message = "apikey 不应该为空"
		c.Data["json"] = setApiKeyResponse
		c.ServeJSON()
		return
	}
	userId := c.GetSession("user_id")
	if userId == nil {
		c.Ctx.Output.SetStatus(401)
		setApiKeyResponse.Code = 1
		setApiKeyResponse.Message = "用户未登录"
		c.Data["json"] = setApiKeyResponse
		c.ServeJSON()
		return
	}
	user, err := models.GetUser(userId.(int))
	if err != nil {
		c.Ctx.Output.SetStatus(401)
		setApiKeyResponse.Code = 1
		setApiKeyResponse.Message = "用户未登录"
		c.Data["json"] = setApiKeyResponse
		c.ServeJSON()
		return
	}
	_, err = models.SetUserApiKey(user.Id, apiKey)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		setApiKeyResponse.Code = 1
		setApiKeyResponse.Message = err.Error()
		c.Data["json"] = setApiKeyResponse
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	setApiKeyResponse.Code = 0
	setApiKeyResponse.Message = "success"
	c.Data["json"] = setApiKeyResponse
	c.ServeJSON()
}

type IsApiKeySetResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	IsSet   bool   `json:"is_set"`
}

func (c *GptController) IsApiKeySet() {
	var isApiKeySetResponse IsApiKeySetResponse
	userId := c.GetSession("user_id")
	if userId == nil {
		c.Ctx.Output.SetStatus(401)
		isApiKeySetResponse.Code = 1
		isApiKeySetResponse.IsSet = false
		isApiKeySetResponse.Message = "用户未登录"
		c.Data["json"] = isApiKeySetResponse
		c.ServeJSON()
		return
	}
	user, err := models.GetUser(userId.(int))
	if err != nil {
		c.Ctx.Output.SetStatus(401)
		isApiKeySetResponse.Code = 1
		isApiKeySetResponse.IsSet = false
		isApiKeySetResponse.Message = "用户未登录"
		c.Data["json"] = isApiKeySetResponse
		c.ServeJSON()
		return
	}
	apikey, err := models.GetApiKeyByUserId(user.Id)
	if apikey == "" {
		c.Ctx.Output.SetStatus(200)
		isApiKeySetResponse.Code = 0
		isApiKeySetResponse.IsSet = false
		isApiKeySetResponse.Message = "success"
		c.Data["json"] = isApiKeySetResponse
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	isApiKeySetResponse.Code = 0
	isApiKeySetResponse.IsSet = true
	isApiKeySetResponse.Message = "success"
	c.Data["json"] = isApiKeySetResponse
	c.ServeJSON()
}

type ChatRequest struct {
	Prompt string `json:"prompt"`
}

type ChatResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    string `json:"data"`
}

func (c *GptController) Chat() {
	var chatRequest ChatRequest
	var chatResponse ChatResponse
	json.NewDecoder(c.Ctx.Request.Body).Decode(&chatRequest)
	prompt := chatRequest.Prompt
	userId := c.GetSession("user_id")
	if userId == nil {
		c.Ctx.Output.SetStatus(401)
		chatResponse.Code = 1
		chatResponse.Message = "用户未登录"
		c.Data["json"] = chatResponse
		c.ServeJSON()
		return
	}
	user, err := models.GetUser(userId.(int))
	if err != nil {
		c.Ctx.Output.SetStatus(401)
		chatResponse.Code = 1
		chatResponse.Message = "用户未登录"
		c.Data["json"] = chatResponse
		c.ServeJSON()
		return
	}
	apikey, err := models.GetApiKeyByUserId(user.Id)
	if apikey == "" {
		c.Ctx.Output.SetStatus(401)
		chatResponse.Code = 1
		chatResponse.Message = "用户未设置apikey"
		c.Data["json"] = chatResponse
		c.ServeJSON()
		return
	}

	chatContext, _ := chatContextMap[string(user.Id)]

	generatedText, err := chatContext.RequestGptService(prompt, apikey, 0.7)
	if err != nil {
		c.Ctx.Output.SetStatus(401)
		chatResponse.Code = 1
		chatResponse.Message = "gpt服务出错" + err.Error()
		c.Data["json"] = chatResponse
		c.ServeJSON()
		return
	}
	chatContextMap[string(user.Id)] = chatContext

	c.Ctx.Output.SetStatus(200)
	chatResponse.Code = 0
	chatResponse.Message = "success"
	chatResponse.Data = generatedText
	c.Data["json"] = chatResponse
	c.ServeJSON()
	return
}

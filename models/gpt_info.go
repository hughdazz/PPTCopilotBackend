package models

import (
	"backend/conf"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type ChatContext struct {
	messages []map[string]string
	mux      sync.Mutex
}

// 用户gpt信息
type GptInfo struct {
	Id     int
	UserId int
	ApiKey string `orm:"size(100)"`
}

func GetApiKeyByUserId(userId int) (string, error) {
	o := orm.NewOrm()
	gptContext := GptInfo{UserId: userId}
	err := o.Read(&gptContext, "UserId")
	if err == orm.ErrNoRows {
		return "", err
	}
	return gptContext.ApiKey, nil
}

func SetUserApiKey(userId int, apiKey string) (string, error) {
	o := orm.NewOrm()
	gptContext := GptInfo{UserId: userId}
	err := o.Read(&gptContext, "UserId")
	if err == orm.ErrNoRows {
		gptContext := GptInfo{UserId: userId, ApiKey: apiKey}
		_, err := o.Insert(&gptContext)
		return apiKey, err
	} else {
		gptContext.ApiKey = apiKey
		_, err := o.Update(&gptContext)
		return apiKey, err
	}
}

func (ctx *ChatContext) AddMessage(role, content string) {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	ctx.messages = append(ctx.messages, map[string]string{"role": role, "content": content})
}

func (ctx *ChatContext) GetContext() []map[string]string {
	ctx.mux.Lock()
	defer ctx.mux.Unlock()
	return ctx.messages
}

func (ctx *ChatContext) RequestGptService(userMessage string, apiKey string, temperature float64) (string, error) {
	ctx.AddMessage("user", userMessage)

	headers := map[string][]string{
		"Authorization": {fmt.Sprintf("Bearer %s", apiKey)},
		"Content-Type":  {"application/json"},
	}

	data := map[string]interface{}{
		"model":       conf.GetGptModel(),
		"messages":    ctx.GetContext(),
		"temperature": temperature,
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	proxyURL, _ := url.Parse(conf.GetGptProxy()) // 使用代理，clash，建议用美国节点
	client := &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}}
	req, err := http.NewRequest("POST", conf.GetGptApiUrl(), bytes.NewBuffer(dataBytes))
	if err != nil {
		return "", err
	}

	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return "", err
	}

	choices := result["choices"].([]interface{})
	generatedText := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	ctx.AddMessage("assistant", generatedText)

	return generatedText, nil
}

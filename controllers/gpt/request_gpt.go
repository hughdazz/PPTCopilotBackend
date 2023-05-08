package gpt

import (
	"backend/conf"
	"fmt"

	"github.com/imroc/req"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}
type ResponseBody struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func RequestGpt(prompt string) (string, error) {
	apikey := conf.GetGptApiKey()

	var body RequestBody
	body.Model = conf.GetGptModel()
	body.Messages = append(body.Messages, Message{
		Role:    "user",
		Content: prompt,
	})
	body.Temperature = 0.7

	req_url := conf.GetGptApiUrl()
	req.SetProxyUrl("http://host.docker.internal:7890")
	// 进行http请求
	resp, err := req.Post(req_url, req.BodyJSON(&body), req.Header(map[string]string{
		"Authorization": "Bearer " + apikey,
		"Content-Type":  "application/json",
	}))
	if err != nil {
		return "", err
	}
	if resp.Response().StatusCode != 200 {
		return "", fmt.Errorf("GPT请求失败，状态码为%d", resp.Response().StatusCode)
	}

	var res ResponseBody
	resp.ToJSON(&res)
	return res.Choices[0].Message.Content, nil
}

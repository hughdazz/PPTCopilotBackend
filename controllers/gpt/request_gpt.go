package gpt

import (
	"backend/conf"
	"fmt"
	"regexp"
	"strings"
	"unicode"

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

func ErrorScanner(gptResponse string) (string, error) {
	// 确保程序以 '<' 开头，以 '>'结尾
	if !(strings.HasPrefix(gptResponse, "<") && strings.HasSuffix(gptResponse, ">")) {
		start := strings.Index(gptResponse, "<")
		end := strings.LastIndex(gptResponse, ">")
		if start == -1 || end == -1 {
			return "", fmt.Errorf("response does not contain required '<' or '>'")
		}
		gptResponse = gptResponse[start : end+1]
	}

	// 移除所有转义字符
	r := regexp.MustCompile(`\\.`)
	gptResponse = r.ReplaceAllString(gptResponse, "")

	// 移除所有不必要空格，"<>"标签内内容除外
	inScope := false
	var builder strings.Builder
	for _, ch := range gptResponse {
		if ch == '<' {
			inScope = true
			builder.WriteRune(ch)
			continue
		}
		if ch == '>' {
			inScope = false
			builder.WriteRune(ch)
			continue
		}
		if !inScope && unicode.IsSpace(ch) {
			continue
		}
		builder.WriteRune(ch)
	}

	return builder.String(), nil
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
	//req.SetProxyUrl("http://host.docker.internal:7890")

	// 最多尝试3次
	retryCount := 0

	for retryCount < 3 {
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

		// 扫描错误
		result, err := ErrorScanner(res.Choices[0].Message.Content)
		if err != nil {
			retryCount++
			continue
		}

		// 成功
		return result, nil
	}

	// 3次尝试均失败
	return "", fmt.Errorf("all retries failed")
}

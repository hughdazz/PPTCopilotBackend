package conf

import (
	"gopkg.in/ini.v1"
)

type GptConfig struct {
	GptApiUrl                  string
	GptModel                   string
	GptProxy                   string
	CatalogPromptTemplate      string
	UpdateSinglePromptTemplate string
}

var GptConfigInstance GptConfig

func init() {
	cfg, err := ini.Load("./conf/gpt.conf")
	if err != nil {
		panic("Failed to read gpt config file: " + err.Error())
	}

	// 读取配置项
	GptConfigInstance = GptConfig{
		GptApiUrl:                  cfg.Section("").Key("gpt_api_url").String(),
		GptModel:                   cfg.Section("").Key("gpt_model").String(),
		GptProxy:                   cfg.Section("").Key("gpt_proxy").String(),
		CatalogPromptTemplate:      cfg.Section("").Key("catalog_prompt_template").String(),
		UpdateSinglePromptTemplate: cfg.Section("").Key("single_page_prompt_template").String(),
	}
}

func GetGptApiUrl() string {
	return GptConfigInstance.GptApiUrl
}

func GetGptModel() string {
	return GptConfigInstance.GptModel
}

func GetGptProxy() string {
	return GptConfigInstance.GptProxy
}

func GetCatalogPromptTemplate() string {
	return GptConfigInstance.CatalogPromptTemplate
}

func GetUpdateSinglePromptTemplate() string {
	return GptConfigInstance.UpdateSinglePromptTemplate
}

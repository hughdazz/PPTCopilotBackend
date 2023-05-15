package conf

import "gopkg.in/ini.v1"

type GptConfig struct {
	GptApiUrl                  string
	GptModel                   string
	GptProxy                   string
	GptApiKey                  string
	GptApiKeys                 []string
	OutlinePromptTemplate      string
	GuideSinglePromptTemplate  string
	UpdateSinglePromptTemplate string
}

type EmailVerifyConfig struct {
	mailer   string
	host     string
	port     int
	username string
	password string
	from     string
}

var GptConfigInstance GptConfig
var EmailVerifyConfigInstance EmailVerifyConfig

func init() {
	cfg, err := ini.Load("./conf/gpt.conf")
	if err != nil {
		panic("Failed to read gpt config file: " + err.Error())
	}

	emailcfg, err := ini.Load("./conf/emailverify.conf")

	if err != nil {
		panic("Failed to read email config file: " + err.Error())
	}

	// 读取配置项
	GptConfigInstance = GptConfig{
		GptApiUrl:                  cfg.Section("").Key("gpt_api_url").String(),
		GptModel:                   cfg.Section("").Key("gpt_model").String(),
		GptProxy:                   cfg.Section("").Key("gpt_proxy").String(),
		GptApiKey:                  cfg.Section("").Key("gpt_api_key").String(),
		GptApiKeys:                 cfg.Section("").Key("gpt_api_keys").Strings(","),
		OutlinePromptTemplate:      cfg.Section("").Key("outline_prompt_template").String(),
		GuideSinglePromptTemplate:  cfg.Section("").Key("guide_single_prompt_template").String(),
		UpdateSinglePromptTemplate: cfg.Section("").Key("single_page_prompt_template").String(),
	}

	EmailVerifyConfigInstance = EmailVerifyConfig{
		mailer:   emailcfg.Section("").Key("mailer").String(),
		host:     emailcfg.Section("").Key("host").String(),
		port:     emailcfg.Section("").Key("port").MustInt(),
		username: emailcfg.Section("").Key("username").String(),
		password: emailcfg.Section("").Key("password").String(),
		from:     emailcfg.Section("").Key("from").String(),
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

func GetGptApiKey() string {
	return GptConfigInstance.GptApiKey
}

func GetGptApiKeys() []string {
	return GptConfigInstance.GptApiKeys
}

func GetOutlinePromptTemplate() string {
	return GptConfigInstance.OutlinePromptTemplate
}

func GetUpdateSinglePromptTemplate() string {
	return GptConfigInstance.UpdateSinglePromptTemplate
}

func GetGuideSinglePromptTemplate() string {
	return GptConfigInstance.GuideSinglePromptTemplate
}

func GetMailer() string {
	return EmailVerifyConfigInstance.mailer
}

func GetHost() string {
	return EmailVerifyConfigInstance.host
}

func GetPort() int {
	return EmailVerifyConfigInstance.port
}

func GetUsername() string {
	return EmailVerifyConfigInstance.username
}

func GetPassword() string {
	return EmailVerifyConfigInstance.password
}

func GetFrom() string {
	return EmailVerifyConfigInstance.from
}

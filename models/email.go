package models

import (
	"backend/conf"
	"context"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/utils"
)

var bm cache.Cache

func init() {
	// 初始化全局缓存变量
	bm, _ = cache.NewCache("memory", `{"interval":300}`)

}

func SendEmail(email string, verificationCode string) error {
	// 发送邮件

	// 从配置文件中获取邮件配置
	username := conf.GetUsername()
	password := conf.GetPassword()
	host := conf.GetHost()
	port := conf.GetPort()
	from := conf.GetFrom()

	config := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%d}`, username, password, host, port)

	verifyemail := utils.NewEMail(config)

	verifyemail.To = []string{email}
	verifyemail.From = from
	verifyemail.Subject = "beego-验证邮件"
	verifyemail.Text = "邮件正文"
	verifyemail.HTML = verificationCode
	err := verifyemail.Send()
	return err

}

func SetCodeCache(email string, verifyCode string) error {
	// 设置验证码缓存

	err := bm.Put(context.TODO(), email, verifyCode, time.Second*300)
	if err != nil {
		return err
	}

	return nil

}

func GetVerifyCode(email string) (interface{}, error) {
	// 检查验证码

	verifyCode, err := bm.Get(context.TODO(), email)
	if verifyCode == nil {
		return "", err
	}
	return verifyCode, nil

}

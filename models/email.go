package models

import (
	"backend/conf"
	"context"
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/utils"
)

type EmailVerification struct {
	Id        int
	Email     string    `orm:"size(100)"`
	Code      string    `orm:"size(100)"`
	ExpiresAt time.Time `orm:"type(datetime)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

var bm cache.Cache

func init() {
	// 初始化全局缓存变量
	bm, _ = cache.NewCache("memory", `{"interval":300}`)

}

func SetVerifyCode(email string, verifyCode string) error {
	o := orm.NewOrm()
	expiresAt := time.Now().Add(time.Hour)
	// Check if record with given email exists
	var emailVerification EmailVerification
	err := o.QueryTable("email_verification").Filter("email", email).One(&emailVerification)
	if err == nil {
		// Record exists, update Code and ExpiresAt fields
		emailVerification.Code = verifyCode
		emailVerification.ExpiresAt = expiresAt
		_, err = o.Update(&emailVerification)
		return err
	} else {
		// Record does not exist, create new record
		emailVerification = EmailVerification{
			Email:     email,
			Code:      verifyCode,
			ExpiresAt: expiresAt,
		}
		_, err = o.Insert(&emailVerification)
		return err
	}
}

func CheckVerifyCode(email string, verifyCode string) (string, error) {
	// Retrieve the EmailVerification row for the given email
	emailVerification := EmailVerification{Email: email}
	if err := orm.NewOrm().Read(&emailVerification, "Email"); err != nil {
		return "", err
	}

	// Check if the verification code is correct and not expired
	if emailVerification.Code != verifyCode {
		return "invalid verification code", nil
	}

	if time.Now().Before(emailVerification.ExpiresAt) {
		return "", nil
	} else {
		return "verification code expired", nil
	}
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

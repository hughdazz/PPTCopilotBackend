package controllers

import (
	"backend/conf"
	"backend/models"
	"encoding/json"

	"fmt"

	"github.com/beego/beego/v2/core/utils"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/google/uuid"
)

type ForgetPasswdController struct {
	beego.Controller
}

type ForgetPasswdRequest struct {
	Email string `json:"email"`
}

type ForgetPasswdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (this *ForgetPasswdController) SendVerificationEmail() {
	var forgetresponse ForgetPasswdResponse

	// 解析json数据
	var forget_request ForgetPasswdRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&forget_request)
	email := forget_request.Email

	//数据库中检查邮箱是否存在
	_, err := models.GetUserByEmail(email)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		forgetresponse.Code = 1
		forgetresponse.Message = "邮箱不存在"
		this.Data["json"] = forgetresponse
		this.ServeJSON()
	}

	// 生成验证码
	verificationCode := uuid.New().String()

	fmt.Println(verificationCode)

	// TODO: 将邮箱地址和验证码保存到数据库中
	err = models.SetVerifyCode(email, verificationCode)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		forgetresponse.Code = 2
		forgetresponse.Message = "储存验证码失败"
		this.Data["json"] = forgetresponse
		this.ServeJSON()
	}

	// 发送验证邮件
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
	err = verifyemail.Send()
	if err != nil {
		// 邮件发送失败
		this.Data["json"] = map[string]interface{}{"success": false, "message": "邮件发送失败"}
		this.ServeJSON()
		return
	}

	// 邮件发送成功
	this.Data["json"] = map[string]interface{}{"success": true, "message": "邮件发送成功"}
	this.ServeJSON()
}

type CheckVerificationEmailRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (this *ForgetPasswdController) CheckVerificationEmail() {
	var checkrequest CheckVerificationEmailRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&checkrequest)
	email := checkrequest.Email
	code := checkrequest.Code

	returncode, err := models.CheckVerifyCode(email, code)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"success": false, "message": "未知的邮箱"}
		this.ServeJSON()
		return
	}

	if returncode != "" {
		this.Data["json"] = map[string]interface{}{"success": false, "message": returncode}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"success": true, "message": "验证成功"}
	this.ServeJSON()
}

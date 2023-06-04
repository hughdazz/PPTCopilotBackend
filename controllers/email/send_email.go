package email

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"

	"github.com/google/uuid"
)

type SendEmailRequest struct {
	Email string `json:"email"`
}

func (this *Controller) SendEmail() {
	var request SendEmailRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)
	email := request.Email

	// _, err := models.GetUserByEmail(email)
	// if err != nil {
	// 	this.Data["json"] = controllers.MakeResponse(controllers.Err, "邮箱不存在", email)
	// 	this.ServeJSON()
	// 	return
	// }

	// 生成验证码
	verificationCode := uuid.New().String()

	// 将验证码保存到缓存中

	err := models.SetCodeCache(email, verificationCode)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "验证码缓存失败", nil)
		this.ServeJSON()
		return
	}

	// 发送邮件
	err = models.SendEmail(email, verificationCode)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "邮件发送失败", email)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "邮件发送成功", nil)
	this.ServeJSON()

}

package email

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
)

type VerifyEmailRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (this *Controller) VerifyEmail() {
	var request VerifyEmailRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)
	email := request.Email
	code := request.Code
	verifycode, err := models.GetVerifyCode(email)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "验证码获取失败", nil)
		this.ServeJSON()
		return
	}

	if code != verifycode {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "验证码错误", verifycode)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "验证成功", nil)
	this.ServeJSON()
}

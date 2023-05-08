package user

import (
	"backend/controllers"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
)

type UploadAvatarRequest struct {
	Avatar string `json:"avatar"`
}

func (this *Controller) UploadAvatar() {
	id_ := this.Ctx.Input.Param(":id")

	var request UploadAvatarRequest
	json.NewDecoder(this.Ctx.Request.Body).Decode(&request)

	// 存在static/user/:id/avatar.png
	// 递归创建目录
	saveDir := "static/user/" + id_
	err := os.MkdirAll(saveDir, 0755)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	// 保存文件
	savePath := saveDir + "/avatar.png"
	// 解码base64
	avatar, err := base64.StdEncoding.DecodeString(request.Avatar)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}
	// 写入文件
	err = ioutil.WriteFile(savePath, avatar, 0644)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), nil)
		this.ServeJSON()
		return
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "Avatar updated", nil)
	this.ServeJSON()
}

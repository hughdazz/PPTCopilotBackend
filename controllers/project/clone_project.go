package project

import (
	"backend/controllers"
	"backend/models"
	"os"
	"strconv"
)

func (this *Controller) CloneProject() {
	// 获取参数
	id_ := this.Ctx.Input.Param(":id")
	clone_id, err := strconv.Atoi(id_)

	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "参数错误", nil)
		this.ServeJSON()
		return
	}

	token, err := this.Ctx.Request.Cookie("token")
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, "未登录", nil)
		this.ServeJSON()
		return
	}
	user_id := models.GetUserId(token.Value)

	// 获取项目信息
	clone_project, err := models.GetProject(clone_id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), 1)
		this.ServeJSON()
		return
	}

	// 获取项目文件
	clone_files, err := models.GetFiles(clone_id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), 2)
		this.ServeJSON()
		return
	}

	// 创建项目
	project, err := models.CreateProject(clone_project.Name, clone_project.Description, user_id)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), 3)
		this.ServeJSON()
		return
	}

	// 创建项目文件夹
	save_dir := models.GetSaveDir(project.Id)
	err = os.MkdirAll(save_dir, 0777)
	if err != nil {
		this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), 4)
		this.ServeJSON()
		return
	}

	// 创建项目文件
	for _, file := range clone_files {
		file_path := save_dir + "/" + file.Name
		old_file_path := models.GetSaveDir(clone_id) + "/" + file.Name
		//将原文件复制入新地址
		err = models.CopyFile(old_file_path, file_path)

		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), 5)
			this.ServeJSON()
			return
		}
		_, err = models.CreateFile(file.Name, project.Id)
		if err != nil {
			this.Data["json"] = controllers.MakeResponse(controllers.Err, err.Error(), 6)
			this.ServeJSON()
			return
		}
	}

	this.Data["json"] = controllers.MakeResponse(controllers.OK, "success", project)
	this.ServeJSON()

}

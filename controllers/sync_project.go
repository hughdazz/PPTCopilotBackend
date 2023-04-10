package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type SyncProjectController struct {
	beego.Controller
}

type SyncProjectRequest struct {
	Project_id int `json:"project_id"`
}

type SyncProjectResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

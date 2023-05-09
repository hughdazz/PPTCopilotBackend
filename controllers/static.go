package controllers

import (
	"net/http"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type StaticRouter struct {
	beego.Controller
}

func (c *StaticRouter) Get() {
	url := c.Ctx.Request.URL.Path
	if strings.Contains(url, "user/") && strings.Contains(url, "avatar.png") {
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "static/user/default.png")
		return
	}
	if strings.Contains(url, "project/") && strings.Contains(url, "cover.png") {
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "static/project/default.png")
		return
	}
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "."+url)
}

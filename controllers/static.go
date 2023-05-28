package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type StaticRouter struct {
	beego.Controller
}

// Get static file
// @router /_static/* [get]
func (c *StaticRouter) Get() {
	url := c.Ctx.Request.URL.Path
	// & /_static/user/:id/avatar.png
	url = url[2:]
	_, err := os.Stat(url)
	if err == nil {
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, url)
		return
	}
	if strings.Contains(url, "user/") && strings.Contains(url, "avatar.png") {
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "static/user/default.png")
		return
	}
	if strings.Contains(url, "project/") && strings.Contains(url, "cover.png") {
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "static/project/default.png")
		return
	}
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, url)
}

// Post static file
// @router /_static/* [post]
func (c *StaticRouter) Post() {
	url := c.Ctx.Request.URL.Path
	url = url[2:] // 去掉前面的 '/_'

	// 确定文件保存路径
	savePath := url

	// 创建文件
	out, err := os.Create(savePath)
	if err != nil {
		c.Data["json"] = MakeResponse(Err, "文件创建失败", nil)
		c.ServeJSON()
		return
	}
	defer out.Close()

	// 从请求主体读取文件数据并写入到文件系统
	_, err = io.Copy(out, c.Ctx.Request.Body)
	if err != nil {
		c.Data["json"] = MakeResponse(Err, "文件写入失败", nil)
		c.ServeJSON()
		return
	}

	// 返回成功消息
	c.Data["json"] = MakeResponse(OK, "文件保存成功", nil)
	c.ServeJSON()
}

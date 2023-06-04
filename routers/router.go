package routers

import (
	"backend/controllers"
	"backend/controllers/email"
	"backend/controllers/gpt"
	"backend/controllers/project"
	"backend/controllers/template"
	"backend/controllers/user"
	"backend/controllers/websocket"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 跨域以及权限控制
	beego.InsertFilter("*", beego.BeforeRouter, cors_access)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/initial", &controllers.InitialController{})
	beego.Router("/_static/*", &controllers.StaticRouter{})
	beego.Router("/upload", &controllers.UploadRouter{})

	// User相关
	userController := beego.NewNamespace("/user",
		beego.NSRouter("/", &user.Controller{}, "get:GetAll;post:CreateUser"),
		beego.NSRouter("/:id", &user.Controller{}, "get:GetUser;put:UpdateUser;delete:DeleteUser"),
		beego.NSRouter("/:id/avatar", &user.Controller{}, "post:UploadAvatar"),
		beego.NSRouter("/:id/project", &user.Controller{}, "get:GetProjects"),
		beego.NSRouter("/:id/favorite", &user.Controller{}, "get:GetFavorites"),

		beego.NSRouter("/login", &user.Controller{}, "post:Login"),
		beego.NSRouter("/logout", &user.Controller{}, "post:Logout"),
	)
	beego.AddNamespace(userController)

	// Project相关
	projectController := beego.NewNamespace("/project",
		beego.NSRouter("/", &project.Controller{}, "get:GetAll;post:CreateProject"),
		beego.NSRouter("/:id", &project.Controller{}, "get:GetProject;post:CloneProject;put:UpdateProject;delete:DeleteProject"),
		beego.NSRouter("/:id/star", &project.Controller{}, "post:StarProject;delete:UnstarProject;get:GetStar"),
		beego.NSRouter("/:id/file", &project.Controller{}, "get:GetFiles;post:CreateFile"),
		beego.NSRouter("/:id/file/update_name", &project.Controller{}, "post:UpdateFileName"),
		beego.NSRouter("/:id/file/:file_name", &project.Controller{}, "get:GetFile;delete:DeleteFile"),
		beego.NSRouter("/:id/json_file", &project.Controller{}, "post:CreateJsonFile"),
		beego.NSRouter("/search", &project.Controller{}, "get:SearchProject"),
	)
	beego.AddNamespace(projectController)

	// Gpt相关
	gptController := beego.NewNamespace("/gpt",
		beego.NSRouter("gen_outline", &gpt.Controller{}, "post:GenOutline"),
		beego.NSRouter("update_slide", &gpt.Controller{}, "post:UpdateSlide"),
		beego.NSRouter("guide_slide", &gpt.Controller{}, "post:GuideSlide"),
		beego.NSRouter("/outline/:id", &gpt.Controller{}, "get:GetOutline;post:UpdateOutline"),
		beego.NSRouter("/gen_ppt", &gpt.Controller{}, "post:GenPPT"),
	)
	beego.AddNamespace(gptController)

	// Template相关
	templateController := beego.NewNamespace("/template",
		beego.NSRouter("/", &template.Controller{}, "get:GetAllTemplates;post:CreateTemplate"),
		beego.NSRouter("/info", &template.Controller{}, "get:GetTemplatesInfo"),
		beego.NSRouter("/:id", &template.Controller{}, "get:GetTemplate"),
		beego.NSRouter("/:id/image", &template.Controller{}, "post:UploadImage"),
	)
	beego.AddNamespace(templateController)

	// email相关
	emailController := beego.NewNamespace("/email",
		beego.NSRouter("send_email", &email.Controller{}, "post:SendEmail"),
		beego.NSRouter("verify_email", &email.Controller{}, "post:VerifyEmail"),
	)
	beego.AddNamespace(emailController)

	// websocket相关
	websocketController := beego.NewNamespace("/ws",
		beego.NSRouter("/", &websocket.Controller{}, "get:Join"),
	)
	beego.AddNamespace(websocketController)

}

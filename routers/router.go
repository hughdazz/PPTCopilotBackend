package routers

import (
	"backend/controllers"
	"backend/models"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/server/web/context"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context) {
		context.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:9529")
		context.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, token")
		context.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")

		// OPTIONS请求直接返回
		if context.Request.Method == "OPTIONS" {
			context.ResponseWriter.WriteHeader(200)
			// 设置跨域
			return
		}
		if context.Request.RequestURI != "/login" && context.Request.RequestURI != "/register" {
			cookie, err := context.Request.Cookie("token")

			if err != nil {
				//进入/projects和/logout必须要有token和cookie
				if context.Request.RequestURI == "/projects" || context.Request.RequestURI == "/logout" {
					context.ResponseWriter.Write([]byte("请先登录"))
					context.Redirect(302, "/login")
				} else if context.Request.Method != "GET" {
					//其他页面有:project_id可以看，但不能改动
					user_id := models.CheckToken(cookie.Value)
					//路由中包含project_id
					if strings.Index(context.Request.RequestURI, "/projects/") >= 0 {
						projectIdStr := context.Input.Param(":project_id")
						projectId, _ := strconv.Atoi(projectIdStr)
						//根据project_id找见项目
						project, pro_err := models.GetProject(projectId)
						//路由错误
						if pro_err != nil {
							context.ResponseWriter.Write([]byte("路由错误"))
							context.Redirect(404, "/login")
						} else if project.Creator.Id != user_id {
							context.ResponseWriter.Write([]byte("您无权操作，请先登录"))
							context.Redirect(302, "/login")
						}
					}

				}
			}
		}
	})

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/projects", &controllers.ProjectsController{})
	beego.Router("/projects/:project_id", &controllers.ProjectController{})
	beego.Router("/projects/:project_id/files", &controllers.FilesController{})
	beego.Router("/projects/:project_id/files/:file_name", &controllers.FileController{})
	beego.Router("/resetpassword", &controllers.ResetPasswordController{})
	beego.Router("/sendverifyemail", &controllers.ForgetPasswdController{}, "post:SendVerificationEmail")
	beego.Router("/checkverifyemail", &controllers.ForgetPasswdController{}, "post:CheckVerificationEmail")
	beego.Router("/checklogin", &controllers.CheckLoginController{})

	gptService := beego.NewNamespace("/gpt",
		beego.NSRouter("/set_api_key", &controllers.GptController{}, "post:SetApiKey"),
		beego.NSRouter("/is_api_key_set", &controllers.GptController{}, "get:IsApiKeySet"),
		beego.NSRouter("/get_catalog", &controllers.GptController{}, "post:GetCatalog"),
		beego.NSRouter("/update_slides", &controllers.GptController{}, "post:UpdateSides"),
		beego.NSRouter("/chat", &controllers.GptController{}, "post:Chat"),
	)
	beego.AddNamespace(gptService)

}

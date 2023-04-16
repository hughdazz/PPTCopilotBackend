package routers

import (
	"backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/projects", &controllers.ProjectsController{})
	beego.Router("/projects/:project_id", &controllers.ProjectController{})
	beego.Router("/projects/:project_id/files", &controllers.FilesController{})
	beego.Router("/resetpassword", &controllers.ResetPasswordController{})
	beego.Router("/sendverifyemail", &controllers.ForgetPasswdController{}, "post:SendVerificationEmail")
	beego.Router("/checkverifyemail", &controllers.ForgetPasswdController{}, "post:CheckVerificationEmail")

	gptService := beego.NewNamespace("/gpt",
		beego.NSRouter("/set_api_key", &controllers.GptController{}, "post:SetApiKey"),
		beego.NSRouter("/is_api_key_set", &controllers.GptController{}, "get:IsApiKeySet"),
		beego.NSRouter("/get_catalog", &controllers.GptController{}, "post:GetCatalog"),
		beego.NSRouter("/update_slides", &controllers.GptController{}, "post:UpdateSides"),
		beego.NSRouter("/chat", &controllers.GptController{}, "post:Chat"),
	)
	beego.AddNamespace(gptService)
}

package routers

import (
	"backend/controllers"
	"backend/controllers/user"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors_access)

	userController := beego.NewNamespace("/user",
		beego.NSRouter("/", &user.Controller{}, "get:GetAllUsers;post:CreateUser"),
		beego.NSRouter("/:id", &user.Controller{}, "get:GetUser;put:UpdateUser;delete:DeleteUser"),
		beego.NSRouter("/:id/password", &user.Controller{}, "put:UpdatePassword"),
		beego.NSRouter("/login", &user.Controller{}, "post:Login"),
		beego.NSRouter("/logout", &user.Controller{}, "post:Logout"),
	)
	beego.AddNamespace(userController)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/projects", &controllers.ProjectsController{})
	beego.Router("/projects/:project_id", &controllers.ProjectController{})
	beego.Router("/projects/:project_id/files", &controllers.FilesController{})
	beego.Router("/projects/:project_id/files/:file_name", &controllers.FileController{})
	beego.Router("/search_project?:filter_words", &controllers.SearchController{})
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

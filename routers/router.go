package routers

import (
	"backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/projects", &controllers.NewProjectController{})
	beego.Router("/register", &controllers.RegisterController{})
}

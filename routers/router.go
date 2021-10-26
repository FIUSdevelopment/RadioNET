package routers

import (
	"radioNet/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/dashboard", &controllers.AdminController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/listen", &controllers.ListenController{})
	beego.Router("/checkusr", &controllers.CheckusrController{})
}

//_ "github.com/mattn/go-sqlite3"
// /usersdb
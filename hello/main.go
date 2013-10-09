package main

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/login", &controllers.LoginController{})
	beego.AutoRouter(&controllers.Login{})
	beego.Run()
}

package routers

import (
	"github.com/astaxie/beego"
	"jxdream/controllers/user"
)

func init() {
	ns := beego.NewNamespace("/user",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&user.UserController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&user.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
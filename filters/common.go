package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
	"regexp"
)

var HasPermission = func(ctx *context.Context) {
	ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
}

/*验证登陆*/
var HasLogin = func(ctx *context.Context) {
	log.Println("requset uri :", ctx.Request.RequestURI)
	log.Println("request data:", string(ctx.Input.RequestBody))
	requestBody := string(ctx.Input.RequestBody)
	if requestBody == "" {

	}

	match, _ := regexp.MatchString("^/user/login/", ctx.Request.RequestURI)

	log.Println("requset uri :", ctx.Request.RequestURI)
	log.Println("whether login :", "OK")
	ok := true
	if !ok && !match {
		ctx.Redirect(302, "/user/login/session/create")
	}
}


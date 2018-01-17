package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
	"regexp"
	"jxdream/libs"
	"jxdream/common"
)

var HasPermission = func(ctx *context.Context) {
	ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
}

/*验证登陆*/
var HasLogin = func(ctx *context.Context) {
	log.Println("JWT token:", string(ctx.Request.Header.Get("Authorization")))

	common.SetParamDate(ctx, )
	log.Println("requset uri :", ctx.Request.RequestURI)
	log.Println("whether login :", ok)
	match, _ := regexp.MatchString("^/user/login/", ctx.Request.RequestURI)
	if !ok && !match {
		ctx.Redirect(302, "/user/login/session/create")
	}
}


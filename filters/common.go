package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"log"
	"regexp"
)

var HasPermission = func(ctx *context.Context) {
	ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
}

/*验证登陆*/
var HasLogin = func(ctx *context.Context) {
	log.Println("sessionID :",ctx.Input.CruSession.SessionID())
	_, ok := ctx.Input.Session("isLogin").(bool)
	log.Println("requset uri :",ctx.Request.RequestURI)
	log.Println("whether login :",ok)
	match, _ := regexp.MatchString("^/user/login/", ctx.Request.RequestURI)
	if !ok && !match {
		ctx.Redirect(302,"/user/login/session/create")
	}
}

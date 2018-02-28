package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
	"regexp"
	"jxdream/common"
	"encoding/json"
	"jxdream/libs"
)

var HasPermission = func(ctx *context.Context) {
	ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
}

/*验证登陆*/
var HasLogin = func(ctx *context.Context) {
	log.Println("requset uri :", ctx.Request.RequestURI)
	log.Println("request data:", string(ctx.Input.RequestBody))
	requestBody := ctx.Input.RequestBody
	isLogin := false		//默认未登陆

	//如果没有请求参数，构建一个默认请求参数
	if string(requestBody) == "" {
		requestParam,_ := common.BuildDefaultRequest()
		requestParamStr,err := json.Marshal(requestParam)
		if err != nil {
			libs.CheckError(err)
		}
		ctx.Input.RequestBody = requestParamStr
	} else {

		requestParam := new(common.RequestParam)
		err := json.Unmarshal(requestBody,requestParam)
		mapClaims, err :=libs.GetClaims(requestParam.Header.JWT)
		if err != nil {
			libs.CheckError(err)
		}
		isLogin,_ = mapClaims["isLogin"].(bool)
	}

	log.Println("whether login :", isLogin)

	match, _ := regexp.MatchString("^/user/login/", ctx.Request.RequestURI)

	if !isLogin && !match {
		ctx.Redirect(302, "/user/login/session/create")
	}
}


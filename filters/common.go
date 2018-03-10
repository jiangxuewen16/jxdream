package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
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

	//如果没有请求参数，构建一个默认请求参数
	if string(requestBody) != "" {
		requestParam := new(common.RequestParam)
		err := json.Unmarshal(requestBody,requestParam)
		jwtClaims, err :=libs.GetClaims(requestParam.Header.JWT)
		if err != nil {
			jwtClaims.IsLogin = false
			requestParam, err := common.BuildRequest(jwtClaims,nil, err.Error(), common.TOKEN_VALIDATION_FAILED)
			libs.CheckError(err)

			requestByte,err := json.Marshal(requestParam)
			libs.CheckError(err)

			ctx.Input.RequestBody = requestByte
		}
		//isLogin = jwtClaims.IsLogin
	} else {
		requestParam,_ := common.BuildDefaultRequest(common.REQUEST_NOT_HAS_HEADER)
		requestParamStr,err := json.Marshal(requestParam)
		if err != nil {
			libs.CheckError(err)
		}
		ctx.Input.RequestBody = requestParamStr
	}

	//log.Println("whether login :", isLogin)

	//登录页面
	/*loginUrl := beego.AppConfig.String("LoginUrl")
	match, _ := regexp.MatchString("^" + loginUrl, ctx.Request.RequestURI)
	if (!isLogin && !match) || string(requestBody) == "" {
		requestParam,_ := common.BuildDefaultRequest(common.REQUEST_NOT_HAS_HEADER)
		requestParamStr,err := json.Marshal(requestParam)
		if err != nil {
			libs.CheckError(err)
		}
		ctx.Input.RequestBody = requestParamStr
	}*/
}


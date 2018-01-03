package user

import (
	"jxdream/models/user"
	"jxdream/controllers"
	"log"
	"jxdream/libs"
)

type LoginController struct {
	controllers.BaseController
}

const (
	SecretKey = "welcome to wangshubo's blog"
)

type Token struct {
	Token string `json:"token"`
}

// @router /session/create [post]
func (this *LoginController) LoginIn() {
	user := user.User{}
	if err := this.SetParamDate(&user); err != nil {
		this.AjaxReturn("参数有误", controllers.ERROR_CODE, err, nil)
	}

	/*数据验证*/
	errList := user.ParamValid()
	if errList != nil {
		this.AjaxReturn("参数有误", controllers.ERROR_CODE, errList, nil)
	}

	isFind := user.FindByUNameAndPwd()
	if ( ! isFind) {
		this.AjaxReturn("登陆失败", controllers.ERROR_CODE, nil, nil)
	}

	/*生成JWT*/
	JWTClaims :=libs.JWTClaims{user.Id, user.Username, user.Nickname, user.Avatar}
	jwtClaims := libs.JWTToken{JWTClaims:JWTClaims}
	tokenString, err := jwtClaims.BuildJWT()
	log.Println(tokenString)

	if err != nil {
		this.AjaxReturn("登陆失败", controllers.ERROR_CODE, nil, nil)
		return
	}

	this.AjaxReturn("登陆成功", controllers.SUCCESS_CODE, nil, tokenString)
}

// @router /session/destroy   [get]
func (this *LoginController) Logout() {
	this.DestroySession()
}

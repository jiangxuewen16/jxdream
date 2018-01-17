package user

import (
	"jxdream/controllers"
	"jxdream/libs"
	"jxdream/models/user"
	"log"
)

type LoginController struct {
	controllers.BaseController
}

// @router /session/create [post]
func (this *LoginController) LoginIn() {
	user := user.User{}
	if err := this.SetParamDate(&user); err != nil {
		this.FailureResponser("参数有误", controllers.ERROR_CODE, nil)
	}

	/*数据验证*/
	errList := user.ParamValid()
	if errList != nil {
		this.FailureResponser("参数有误", controllers.ERROR_CODE, errList)
	}

	isFind := user.FindByUNameAndPwd()
	if !isFind {
		this.FailureResponser("登陆失败", controllers.ERROR_CODE, nil)
	}

	/*生成JWT*/
	JWTClaims := libs.JWTClaims{user.Id, user.Username, user.Nickname, user.Avatar}
	tokenString, err := libs.BuildJWT(JWTClaims)
	log.Println(tokenString)

	if err != nil {
		this.FailureResponser("登陆失败", controllers.ERROR_CODE, nil)
		return
	}

	this.SuccessResponser("登陆成功", nil)
}

// @router /session/destroy   [get]
func (this *LoginController) Logout() {
	this.DestroySession()
}

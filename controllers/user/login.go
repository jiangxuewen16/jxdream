package user

import (
	"jxdream/controllers"
	"jxdream/models/user"
	"jxdream/common"
)

type LoginController struct {
	controllers.BaseController
}

// @router /session/create [post]
func (this *LoginController) LoginIn() {
	user := user.User{}
	if err := this.GetDataParam(&user); err != nil {
		this.FailureResponser("参数绑定错误，请检查您的参数", common.ERROR_CODE, nil)
	}

	/*数据验证*/
	if errList := user.ParamValid(); errList != nil {
		this.FailureResponser("参数有误", common.ERROR_CODE, errList)
	}

	isFind := user.FindByUNameAndPwd()
	if !isFind {
		this.FailureResponser("登陆失败", common.ERROR_CODE, nil)
	}

	/*生成JWT*/
	/*JWTClaims := libs.JWTClaims{user.Id, user.Username, user.Nickname, user.Avatar, }
	tokenString, err := libs.BuildJWT(JWTClaims)
	log.Println(tokenString)

	if err != nil {
		this.FailureResponser("登陆失败", common.ERROR_CODE, nil)
		return
	}*/

	this.SetUserBaseInfo(&user)
	this.SuccessResponser("登陆成功", nil)
}

// @router /session/destroy   [get]
func (this *LoginController) Logout() {
	this.DestroySession()
}

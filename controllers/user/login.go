package user

import (
	"jxdream/controllers"
	"jxdream/models/user"
	"jxdream/common"
	"fmt"
	"jxdream/models"
)

type LoginController struct {
	controllers.BaseController
}

// @router /session/create [post]
func (this *LoginController) LoginIn() {
	user := user.User{}
	if err := this.GetDataParam(&user); err != nil {
		this.FailureResponser("参数绑定错误，请检查您的参数", common.REQUEST_PARAMETER_BIND_ERROR, nil)
	}

	a := models.Redis.Get("name")
	models.Redis.Put("age",22,100)
	fmt.Println(a)

	/*数据验证*/
	if errList := user.ParamValid(); errList != nil {
		this.FailureResponser("参数有误", common.REQUEST_PARAMETER_VALID_ERROR, errList)
	}

	isFind := user.FindByUNameAndPwd()
	if !isFind {
		this.FailureResponser("登陆失败", common.USER_LOGIN_FIALD, nil)
	}

	this.SetUserBaseInfo(&user)
	this.SuccessResponser("登陆成功", nil)
}

// @router /session/destroy   [get]
func (this *LoginController) Logout() {
	this.DestroySession()
}

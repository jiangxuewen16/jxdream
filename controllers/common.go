package controllers

import (
	"github.com/astaxie/beego"
	"jxdream/common"
	"jxdream/libs"
	"jxdream/models/user"
	"net/http"
	"strings"
	"encoding/json"
)

type BaseController struct {
	beego.Controller

	UserBaseInfo   libs.JWTClaims //jwt基本信息
	ControllerName string
	ActionName     string
	User           *user.User
	BaseUrl        string
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.ControllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
	this.ActionName = strings.ToLower(actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.BaseUrl = this.Ctx.Request.URL.String()

	headerParam, err := this.GetHeaderParam()
	if err != nil {
		libs.CheckError(err)
	}

	//检验该url是否要检验jwt  todo:不需要检验是否要验证登录，这个可以放到拦截器去
	/*urlStr := beego.AppConfig.String("notCheckLoginUrl")
	if urls := strings.Split(urlStr, ","); len(urls) > 0 {
		if libs.StringArrayHasElement(urls, this.BaseUrl) {
			return
		}
	}*/

	jwtToken := headerParam.JWT
	jwtClaims, err := libs.GetClaims(jwtToken)
	if err != nil {
		libs.CheckError(err)
	}

	libs.BuildJWT(jwtClaims)
	if !jwtClaims.IsLogin {
		resonseParam := &common.ResponseParam{}
		err := common.SetParamDate(this.Ctx, resonseParam)
		libs.CheckError(err)

		this.Data["json"] = resonseParam
		this.ServeJSON()
		this.StopRun()
	}

	this.UserBaseInfo = jwtClaims
	/*this.UserBaseInfo.UserId, _ = mapClaims["userId"].(int)
	this.UserBaseInfo.UserName, _ = mapClaims["userName"].(string)
	this.UserBaseInfo.NickName, _ = mapClaims["nickName"].(string)
	this.UserBaseInfo.Avatar, _ = mapClaims["avatar"].(string)
	this.UserBaseInfo.IsLogin, _ = mapClaims["isLogin"].(bool)*/

	//this.jwtClaims = libs.JWTClaims{this.UserId, this.UserName, this.NickName, this.Avatar, this.IsLogin}

}

//重定向（web）
/*
func (this *BaseController) Redirectd(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}
*/

/*是否是post提交*/
func (this *BaseController) IsPost() bool {
	return this.Ctx.Request.Method == "POST"
}

/*获取请求头*/
func (this *BaseController) GetHeader() http.Header {
	return this.Ctx.Request.Header
}

/*获取请求类型*/
func (this *BaseController) GetContentType() string {
	return this.GetHeader().Get("Content-Type")
}

func (this *BaseController) GetRequestParam() (*common.RequestParam, error) {
	requestParam := new(common.RequestParam)
	err := common.SetParamDate(this.Ctx, requestParam)
	if err != nil {
		return nil,err
	}
	return requestParam,nil
}

/*获取data数据，并绑定结构体*/
func (this *BaseController) GetDataParam(struc interface{}) error {
	requestParam, err := this.GetRequestParam()
	if err != nil {
		return err
	}
	info, err := json.Marshal(requestParam.Data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(info, struc)
	if err != nil {
		return err
	}
	return nil
}

//获取header数据
func (this *BaseController) GetHeaderParam() (*common.Header, error) {
	requestParam, err := this.GetRequestParam()
	if err != nil {
		return nil,err
	}
	return requestParam.Header, nil
}

//返回数据
func (this *BaseController) Responser(data interface{}, message string, code int) {
	jwtClaims := this.UserBaseInfo
	response, err := common.BuildRespose(jwtClaims, data, message, code)
	libs.CheckError(err)

	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()
}

//成功返回
func (this *BaseController) SuccessResponser(message string, data interface{}) {
	this.Responser(data, message, 200)
}

//失败返回
func (this *BaseController) FailureResponser(message string, code int, data interface{}) {
	this.Responser(data, message, code)
}

func (this *BaseController) SetUserBaseInfo(user *user.User) {
	this.UserBaseInfo.UserId = user.Id
	this.UserBaseInfo.UserName = user.Username
	this.UserBaseInfo.UserType = user.UserType
	this.UserBaseInfo.IsLogin = true
	this.UserBaseInfo.Avatar = user.Avatar
	this.UserBaseInfo.NickName = user.Nickname
}

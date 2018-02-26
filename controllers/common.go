package controllers

import (
	"github.com/astaxie/beego"
	"jxdream/common"
	"jxdream/libs"
	"jxdream/models/user"
	"log"
	"net/http"
	"strings"
)

type BaseController struct {
	beego.Controller
	ControllerName string
	ActionName     string
	IsLogin        bool
	User           *user.User
	UserId         int
	NickName       string
	Avatar         string
	UserName       string
	BaseUrl        string
	jwtClaims      libs.JWTClaims //todo:这个地方需要设置值
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.ControllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.ActionName = strings.ToLower(actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.BaseUrl = this.Ctx.Request.URL.String()

	requesrParam := &common.RequestParam{}
	this.SetParamDate(requesrParam)

	urlStr := beego.AppConfig.String("notCheckLoginUrl")
	if urls := strings.Split(urlStr, ","); len(urls) > 0 {
		if libs.StringArrayHasElement(urls, this.BaseUrl) {
			return
		}
	}

	jwtToken := requesrParam.Header.JWT
	mapClaims, err := libs.GetClaims(jwtToken)
	if err != nil {
		log.Println("error:", err)
		this.StopRun()
	}

	this.UserId, _ = mapClaims["userId"].(int)
	this.UserName, _ = mapClaims["userName"].(string)
	this.NickName, _ = mapClaims["nickName"].(string)
	this.Avatar, _ = mapClaims["avatar"].(string)

	this.jwtClaims = libs.JWTClaims{this.UserId, this.UserName, this.NickName, this.Avatar}

}

func (this *BaseController) Redirectd(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

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

/*绑定请求参数到结构体*/
func (this *BaseController) SetParamDate(struc interface{}) error {
	common.SetParamDate(this.Ctx, struc)
	this.ParseForm(struc)
	return nil
}

//返回数据
func (this *BaseController) Responser(data interface{}, message string, code int) {
	jwtClaims := this.jwtClaims
	response, err := common.BuildRespose(jwtClaims, data, message, code)
	if err != nil {
		log.Println("Responser error:", err)
		this.StopRun()
	}

	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) SuccessResponser(message string, data interface{}) {
	this.Responser(data, message, 200)
}

func (this *BaseController) FailureResponser(message string, code int, data interface{}) {
	this.Responser(data, message, code)
}

//获取请求data参数
func (this *BaseController) GetRequstData() interface{} {
	requestParam := &common.RequestParam{}
	this.SetParamDate(requestParam)
	return requestParam.Data
}

//错误panic
//todo:抛错，应该在一个错误页面
/*
func (this *BaseController) ()  {

}
*/

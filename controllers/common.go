package controllers

import (
	"github.com/astaxie/beego"
	"jxdream/common"
	"jxdream/libs"
	"jxdream/models/user"
	"log"
	"net/http"
	"strings"
	"encoding/json"
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

	headerParam := this.GetHeaderParam()

	//检验该url是否要检验jwt  todo:不需要检验是否要验证登录
	/*urlStr := beego.AppConfig.String("notCheckLoginUrl")
	if urls := strings.Split(urlStr, ","); len(urls) > 0 {
		if libs.StringArrayHasElement(urls, this.BaseUrl) {
			return
		}
	}*/

	jwtToken := headerParam.JWT
	mapClaims, err := libs.GetClaims(jwtToken)
	if err != nil {
		log.Println("error:", err)
		this.StopRun()
	}

	this.UserId, _ = mapClaims["userId"].(int)
	this.UserName, _ = mapClaims["userName"].(string)
	this.NickName, _ = mapClaims["nickName"].(string)
	this.Avatar, _ = mapClaims["avatar"].(string)
	this.IsLogin,_ = mapClaims["isLogin"].(bool)

	this.jwtClaims = libs.JWTClaims{this.UserId, this.UserName, this.NickName, this.Avatar, this.IsLogin}

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

/*获取data数据，并绑定结构体*/
func (this *BaseController) GetDataParam(struc interface{}) error {
	requestParam := new(common.RequestParam)
	common.SetParamDate(this.Ctx, requestParam)
	info,_ := json.Marshal(requestParam.Data)
	json.Unmarshal(info, struc)
	return nil
}

//获取header数据
func (this *BaseController) GetHeaderParam() *common.Header {
	requestParam := &common.RequestParam{}
	common.SetParamDate(this.Ctx, requestParam)
	return requestParam.Header
}

//返回数据
func (this *BaseController) Responser(data interface{}, message string, code int) {
	jwtClaims := this.jwtClaims
	response, err := common.BuildRespose(jwtClaims, data, message, code)
	libs.CheckError(err)

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

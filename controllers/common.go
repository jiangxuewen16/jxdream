package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"encoding/json"
	"errors"
	"log"
	"jxdream/models/user"
	"jxdream/common"
	"jxdream/libs"
)

const (
	RJSON      = "application/json"
	RXML       = "application/xml"
	RPLAIN     = "text/plain"
	RHTML      = "text/html"
	RFILE_FORM = "multipart/form-data"
	RFORM      = "application/x-www-form-urlencoded"
)

const (
	ERROR_CODE   = 0
	SUCCESS_CODE = 1
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
	jwtClaims      libs.JWTClaims		//todo:这个地方需要设置值
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.ControllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
	this.ActionName = strings.ToLower(actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.BaseUrl = this.Ctx.Request.URL.String()

	requesrParam := &common.RequestParam{}
	this.SetParamDate(requesrParam)

	jwtToken := requesrParam.Header.JWT
	mapClaims, err := libs.GetClaims(jwtToken)
	if err != nil {
		log.Println("error:", err)
		this.StopRun()
	}

	this.jwtClaims = mapClaims  //todo:??????

	this.UserId, _ = mapClaims["userId"].(int)
	this.UserName, _ = mapClaims["userName"].(string)
	this.NickName, _ = mapClaims["nickName"].(string)
	this.Avatar, _ = mapClaims["avatar"].(string)

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
	requestType := strings.ToLower(this.GetContentType())
	strings.ToLower(this.GetContentType())
	log.Println("requst type :", requestType)
	log.Println("requst params", string(this.Ctx.Input.RequestBody))

	var err error

	switch strings.Split(requestType, ";")[0] {
	case RJSON:
		err = json.Unmarshal(this.Ctx.Input.RequestBody, struc)
	case RFORM, RXML:
		err = this.ParseForm(struc)
	default:
		//TODO:
		err = errors.New("请求类型：" + requestType + "无法解析")
	}
	return err
}

//返回数据
func (this *BaseController) Responser(data interface{}) {
	jwtClaims := this.jwtClaims
	response,err := common.BuildRespose(jwtClaims, data, 200)
	if err != nil {
		log.Println("Responser error:", err)
		this.StopRun()
	}

	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()
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

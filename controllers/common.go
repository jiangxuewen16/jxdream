package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"encoding/json"
	"errors"
	"log"
	"jxdream/models/user"
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
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.ControllerName = strings.ToLower(controllerName[0: len(controllerName)-10])
	this.ActionName = strings.ToLower(actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.BaseUrl = this.Ctx.Request.URL.String()

	//TODO:这里需要改进，不需要session
	this.IsLogin,_ = this.GetSession("isLogin").(bool)
	this.UserId,_ = this.GetSession("userId").(int)
	this.UserName,_ = this.GetSession("userName").(string)
	this.NickName,_ = this.GetSession("nickName").(string)
	this.Avatar,_ = this.GetSession("avatar").(string)

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

	switch strings.Split(requestType,";")[0] {
	case RJSON:
		err = json.Unmarshal(this.Ctx.Input.RequestBody, struc)
	case RFORM,RXML:
		err = this.ParseForm(struc)
	default:
		//TODO:
		err = errors.New("请求类型：" + requestType + "无法解析")
	}
	return err
}

/**
 * 接口返回
 */
func (this *BaseController) AjaxReturn(msg interface{}, code int, data interface{}, JWT interface{}) {
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg
	out["data"] = data
	out["JWT"] = JWT
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}
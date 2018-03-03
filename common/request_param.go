package common

import (
	"jxdream/libs"
	"time"
	"strings"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/context"
	//"github.com/astaxie/beego"
	"log"
)

const (
	RJSON      = "application/json"
	RXML       = "application/xml"
	RPLAIN     = "text/plain"
	RHTML      = "text/html"
	RFILE_FORM = "multipart/form-data"
	RFORM      = "application/x-www-form-urlencoded"
)

type Header struct {
	JWT     string `json:"jwt"`
	Time    int64  `json:"time"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RequestParam struct {
	Header *Header     `json:"header"`
	Data   interface{} `json:"data"`
}

type ResponseParam struct {
	Header *Header     `json:"header"`
	Data   interface{} `json:"data"`
}

func (r *RequestParam) ToString() string {
	RByte, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(RByte)
}

//构建返回数据
func BuildRespose(jwtClaims libs.JWTClaims, data interface{}, message string, code int) (ResponseParam, error) {
	resposeParam := ResponseParam{}
	header, err := buildHeader(jwtClaims, message, code)

	if err != nil {
		return resposeParam, nil
	}

	resposeParam.Header = &header
	resposeParam.Data = data

	return resposeParam, nil
}

//构建请求数据
func BuildRequest(jwtClaims libs.JWTClaims, data interface{}, message string, code int) (RequestParam, error) {
	requestParam := RequestParam{}
	header, err := buildHeader(jwtClaims, message, code)

	if err != nil {
		return requestParam, nil
	}

	requestParam.Header = &header
	requestParam.Data = data

	return requestParam, nil
}

//构建默认请求数据
func BuildDefaultRequest(code int) (requestParam RequestParam, err error) {
	jwtClaims := libs.JWTClaims{IsLogin: false}
	message := "未验证权限"

	requestParam, err = BuildRequest(jwtClaims, nil, message, code)
	if err != nil {
		log.Fatal("jwt 生成失败")
	}

	return requestParam, nil
}

//构建header
func buildHeader(jwtClaims libs.JWTClaims, message string, code int) (Header, error) {
	tokenString, err := libs.BuildJWT(jwtClaims)

	if err != nil {
		return Header{}, err
	}

	header := Header{tokenString, time.Now().Unix(), code, message}
	return header, nil
}

func SetParamDate(ctx *context.Context, struc interface{}) error {
	ctx.Request.Header.Get("Content-Type")
	requestType := strings.ToLower(ctx.Request.Header.Get("Content-Type"))

	var err error

	switch strings.Split(requestType, ";")[0] {
	case RJSON:
		err = json.Unmarshal(ctx.Input.RequestBody, struc)
	case RFORM, RXML:
		//err = beego.ParseForm(struc)  todo:what?????
	default:
		//TODO:
		err = errors.New("请求类型：" + requestType + "无法解析")
	}
	return err
}

package common

import (
	"jxdream/libs"
	"time"
	"strings"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
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

//构建返回数据
func BuildRespose(jwtClaims libs.JWTClaims, data interface{}, message string, code int) (RequestParam, error) {
	requestParam := RequestParam{}
	header, err := buildHeader(jwtClaims, message, code)

	if err != nil {
		return requestParam, nil
	}

	requestParam.Header = &header
	requestParam.Data = data

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
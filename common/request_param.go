package common

import (
	"jxdream/libs"
	"time"
	"encoding/json"
)

type Header struct {
	JWT  string `json:"jwt"`
	Time int64  `json:"time"`
	Code int    `json:"code"`
}

type RequestParam struct {
	Header *Header `json:"header"`
	Data   interface{}   `json:"data"`
}

type ResponseParam struct {
	Header *Header `json:"header"`
	Data   interface{}    `json:"data"`
}

//构建返回数据
func BuildRespose(jwtClaims libs.JWTClaims,data interface{}, code int) (RequestParam, error) {
	requestParam := RequestParam{}
	header,err := buildHeader(jwtClaims, code)

	if err != nil {
		return requestParam, nil
	}

	requestParam.Header = &header
	requestParam.Data = data

	return requestParam,nil
}

//构建header
func buildHeader(jwtClaims libs.JWTClaims, code int) (Header,error) {
	tokenString,err := libs.BuildJWT(jwtClaims)

	if err != nil {
		return Header{},err
	}

	header := Header{tokenString, time.Now().Unix(), code}
	return header,nil
}

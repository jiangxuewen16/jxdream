package common

import (
	"jxdream/libs"
	"time"
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

	header := Header{tokenString, time.Now().Unix(), message, code}
	return header, nil
}

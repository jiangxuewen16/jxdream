package common

import (
	"jxdream/libs"
	"jxdream/models/user"
	"time"
	"encoding/json"
)

type Header struct {
	JWT         string
	RequestTime int64
}

type Data struct {
	Data interface{}
}

type RequestParam struct {
	Header *Header
	Data   *Data
}

type ResponseParam struct {
	Header *Header
	Data   *Data
}

func (this *RequestParam) BuildRespose(data interface{}) (string, error) {



	requestParam := new(RequestParam)
	err := json.Unmarshal([]byte(resquest), requestParam)
	if err != nil {
		return "",err
	}

	requestParam.Header.JWT =

	jwtClaims := libs.JWTClaims{user.Id, user.Username, user.Nickname, user.Avatar}

	this.LoginBuildRespose();

	/*生成JWT*/
	jwtClaims := libs.JWTClaims{user.Id, user.Username, user.Nickname, user.Avatar}
	tokenString, err := jwtClaims.BuildJWT()
	this.Header.JWT =
}

func (this *RequestParam) LoginBuildRespose(jwtClaims libs.JWTClaims, data interface{}) string {
	/*生成JWT*/
	tokenString, err := libs.BuildJWT()
	if err != nil {

	}

	this.Header.JWT = tokenString
	this.Header.RequestTime = time.Now().Unix()
	this.Data.Data = data

	jsonData, err := json.Marshal(this)
	if err != nil {

	}

	return string(jsonData)
}

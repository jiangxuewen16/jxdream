package libs

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
	"github.com/astaxie/beego"
)

type JWTToken struct {
	JWTClaims
	*jwt.Token
}


type JWTClaims struct {
	UserId int
	UserName string
	NickName string
	Avatar string
}

var SecretKey string

func init()  {
	SecretKey = beego.AppConfig.String("SecretKey")			//SecretKey
}

func (this *JWTToken) BuildJWT() (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userId"] = this.UserId
	claims["userName"] = this.UserName
	claims["nickName"] = this.NickName
	claims["avatar"] = this.Avatar
	token.Claims = claims

	//SecretKey := beego.AppConfig.String("SecretKey")		//SecretKey
	tokenString, err = token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Println("生成JWT失败！")
		return
	}
	return
}

func (this *JWTClaims) buildClaims()  {
	this.UserId =
}


func (this *JWTClaims) ValidJWT(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {

	}

	return t,nil
}
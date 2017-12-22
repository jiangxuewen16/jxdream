package libs

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
	"github.com/astaxie/beego"
)

func BuildJWT(id int, username string, nickname string, avatar string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userId"] = id
	claims["userName"] = username
	claims["nickName"] = nickname
	claims["avatar"] = avatar
	token.Claims = claims

	SecretKey := beego.AppConfig.String("SecretKey")		//SecretKey
	tokenString, err = token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Println("生成JWT失败！")
		return

	}
	return
}
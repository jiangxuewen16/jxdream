package libs

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type JWTClaims struct {
	UserId   int
	UserName string
	NickName string
	Avatar   string
	IsLogin  bool
}

type JWTToken struct {
	JWTClaims
	*jwt.Token
}

var SecretKey string

func init() {
	SecretKey = beego.AppConfig.String("SecretKey") //SecretKey
}

//生成jwt token
func BuildJWT(jwtClaims JWTClaims) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := buildClaims(jwtClaims)
	token.Claims = claims

	tokenString, err = token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Println("生成JWT失败！")
		return
	}
	return
}

//构建claims参数
func buildClaims(jwtClaims JWTClaims) jwt.MapClaims {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userId"] = jwtClaims.UserId
	claims["userName"] = jwtClaims.UserName
	claims["nickName"] = jwtClaims.NickName
	claims["avatar"] = jwtClaims.Avatar
	return claims
}

//验证jwt token
func ValidJWT(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	CheckError(err)

	return t, nil
}

func GetClaims(token string) (map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	CheckError(err)
	mapClaims := t.Claims.(jwt.MapClaims)

	return mapClaims, nil
}

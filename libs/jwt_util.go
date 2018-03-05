package libs

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"github.com/goinggo/mapstructure"
)

type JWTClaims struct {
	UserId   int
	UserName string
	UserType int
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
func BuildJWT(jwtClaims JWTClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := buildClaims(jwtClaims)
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Println("生成JWT失败！")
		return "",err
	}
	return tokenString,nil
}

//构建claims参数
func buildClaims(jwtClaims JWTClaims) jwt.MapClaims {
	expireTime,_ := beego.AppConfig.Int("Expire")		//获取配置的超时时间
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userId"] = jwtClaims.UserId
	claims["userName"] = jwtClaims.UserName
	claims["nickName"] = jwtClaims.NickName
	claims["avatar"] = jwtClaims.Avatar
	claims["userType"] = jwtClaims.UserType
	return claims
}

//验证jwt token
func ValidJWT(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	return t, err
}

//获取token信息
func GetClaims(token string) (JWTClaims, error) {
	jwtClaims := JWTClaims{}

	t, err := ValidJWT(token)
	if err != nil {
		return jwtClaims, err
	}

	mapClaims := t.Claims.(jwt.MapClaims)
	mapstructure.Decode(mapClaims,&jwtClaims)
	return jwtClaims, nil
}

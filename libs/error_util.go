package libs

import (
	"log"
	"errors"
)

var (
	ErrAbort = errors.New("User stop run")
)

//输出错误并退出
func CheckError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
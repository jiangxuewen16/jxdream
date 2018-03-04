package common

//请求成功失败code
const (
	ERROR_CODE   = 0
	SUCCESS_CODE = 1

	STATUS_OK = 200

	USER_NOT_LOGIN = 1010

	USER_NOT_PERMISSION = 1020
)

var status_msg = map[int]string{
	STATUS_OK: "OK",

	USER_NOT_LOGIN: "Not Login",

	USER_NOT_PERMISSION: "Not Have Permission",
}

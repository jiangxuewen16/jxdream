package user

import (
	"jxdream/controllers"
	"fmt"
	"jxdream/models"
)

type UserController struct {
	controllers.BaseController
}

// @router /index [get]
func (this *UserController) Index() {
	a := models.Redis.Get("name").([]byte)
	fmt.Println(string(a))
}

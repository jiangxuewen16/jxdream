package user

import (
	"jxdream/controllers"
	"log"
)

type UserController struct {
	controllers.BaseController
}

// @router /index [get]
func (this *UserController) Index() {
	log.Println("whether login :", this.IsLogin, ";", "userId :", this.UserId, ";", "userName :", this.UserName)

}

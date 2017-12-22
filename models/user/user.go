package user

import (
	"github.com/astaxie/beego/validation"
	"fmt"
	"jxdream/libs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int    `orm:"type(int);pk;auto" form:"-"`
	Username   string `orm:"type(varchar);unique;size(32)" form:"username" valid:"Required;MinSize(6);MaxSize(32)"`
	Password   string `orm:"type(char);size(32)" form:"password" valid:"MinSize(6);MaxSize(16)"`
	Mobile     string `orm:"type(varchar);size(15)" form:"mobile"`
	Nickname   string `orm:"type(varchar);size(50)" form:"nickname"`
	Avatar string `orm:"type(varchar);size(50)" form:"avatar"`
	UserType   int    `orm:"type(tinyint)" form:"-"`
	State      int    `orm:"type(tinyint);size(1);default(1)" form:"-"`
	IsDelete   int    `orm:"type(tinyint);size(1);default(0)" form:"-"`
	CreateTime string `orm:"type(datetime)" form:"-"`
	UpdateTime string `orm:"type(timestamp);auto_now_add" form:"-"`
}

//数据验证
func (user *User) ParamValid() (error map[string]*validation.Error) {
	valid := validation.Validation{}
	b, err := valid.Valid(user)
	if err != nil {
		panic(err)
	}

	if !b {
		return valid.ErrorsMap
	}
	return nil
}

func (user *User) FindByUNameAndPwd() bool {
	user.Password = libs.Md5([]byte(user.Password))
	err := orm.NewOrm().Read(user, "username", "password")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
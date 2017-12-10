package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	"errors"
)

var o orm.Ormer //全局orm

func init() {
	conString,err := getDBConf()
	if err != nil {
		panic(err)
	}
	orm.RegisterDataBase("default", "mysql", conString, 30)
	orm.RegisterModelWithPrefix("j_", new(User))
	orm.RunSyncdb("default", false, true)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm() //初始化
}

func getDBConf() (str string, err error) {
	if beego.AppConfig.String("mysqluser") == "" {
		err = errors.New("未配置数据库用户")
		return
	}

	if beego.AppConfig.String("mysqlpass") == "" {
		err = errors.New("未配置数据库密码")
		return
	}

	if beego.AppConfig.String("mysqladdr") == "" {
		err = errors.New("未配置数据库主机")
		return
	}

	if beego.AppConfig.String("mysqlport") == "" {
		err = errors.New("未配置数据库端口")
		return
	}

	if beego.AppConfig.String("mysqldb") == "" {
		err = errors.New("未选择数据库")
		return
	}

	str = beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@(" +
		beego.AppConfig.String("mysqladdr") + ":" + beego.AppConfig.String("mysqlport") + ")/" +
		beego.AppConfig.String("mysqldb") + "?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai"
	fmt.Println("connect sql: ",str)
	return
}

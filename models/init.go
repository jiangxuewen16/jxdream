package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"jxdream/models/user"
	"time"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"log"
)

var Redis cache.Cache

func init() {
	conString, err := getDBConf()
	if err != nil {
		panic(err)
	}
	orm.RegisterDataBase("default", "mysql", conString, 30)

	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterModelWithPrefix("jx_", new(user.User))
	orm.RunSyncdb("default", false, true)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	Redis, err = cache.NewCache("redis",`{"conn":"118.126.105.239:6378", "key":"beecacheRedis", "dbNum":"0", "password":"jiangxuewen"}`)
	if err != nil {
		log.Println(err)
	}
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
	fmt.Println("connect sql: ", str)
	return
}

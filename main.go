package main

import (
	_ "jxdream/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	_ "jxdream/common"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)		//日志

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		orm.Debug = true
	}
	beego.Run()
}

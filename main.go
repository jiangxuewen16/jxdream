package main

import (
	_ "jxdream/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		orm.Debug = true
	}
	beego.Run()
}
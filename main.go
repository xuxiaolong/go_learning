package main

import (
	"github.com/astaxie/beego/orm"
	"myproject/models"
	_ "myproject/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	beego.Run()
}

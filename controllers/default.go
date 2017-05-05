package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["UserName"] = "xuxiaolong"
	type U struct {
		Name string
		Age  int
		Sex  string
	}

	user := &U{
		Name: "xxl",
		Age:  10,
		Sex:  "Male",
	}
	c.Data["User"] = user
	number := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	c.Data["number"] = number
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {

	f, h, err := c.GetFile("uploadname")
	fmt.Println("getfile err:", err)

	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		c.SaveToFile("uploadname", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	}
}

package main

import (
	"PrintYun/DBSource"
	v1 "PrintYun/Router/v1"
	"github.com/kataras/iris/v12"
)


func main() {
	app := CreateApp()
	v1.Service(app)
	app.Run(iris.Addr(":8080"))
}


func CreateApp() *iris.Application{
	DBSource.DBConnect()
	app := iris.Default()
	//yaag.Init(&yaag.Config{
	//	On:       true,                 //是否开启自动生成API文档功能
	//	DocTitle: "Iris",
	//	DocPath:  "ApiList/apidoc.html",        //生成API文档名称存放路径
	//	BaseUrls: map[string]string{"Production": "", "Staging": ""},
	//})
	//app.Use(irisyaag.New()) 		// 开启接口生成文档
	return app
}
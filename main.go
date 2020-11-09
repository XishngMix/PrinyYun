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
	return app
}
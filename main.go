package main

import (
	"PrintYun/Router/v1"
	"github.com/kataras/iris/v12"
	"PrintYun/DBSource"
)

func main() {
	app := CreateApp()

	v1.Service(app)

	app.Run(iris.Addr(":8080"))
}


func CreateApp() *iris.Application{
	app := iris.Default()
	app.Use(myMiddleware)    								// 调用中间件

	DBSource.DBConnect()

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
	return app
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
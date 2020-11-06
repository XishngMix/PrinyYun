package v1

import (
	"PrintYun/Router/v1/View"
	"github.com/kataras/iris/v12"
)

func Service(app *iris.Application)  {
	UserService := app.Party("/user")
	UserService.Post("/sendwxcode", View.WechatLogin)
}
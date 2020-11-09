package v1

import (
	"PrintYun/Router/v1/View"
	"github.com/kataras/iris/v12"
)

func Service(app *iris.Application)  {
	UserService := app.Party("/user")
	UserService.Get("/sendwxcode", View.WechatLogin)
}

//func RouterAuthority(app *iris.Application)  {
//	main := app.Party("/", middleware.CrsAuth()).AllowMethods(iris.MethodOptions)
//	{
//		route := main.Party("/Wechat")
//		{
//			route.Post("/user/sendwxcode", View.WechatLogin).Name = "微信用户登陆"
//		}
//	}
//}
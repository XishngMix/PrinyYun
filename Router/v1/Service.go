package v1

import (
	"PrintYun/Router/v1/Services"
	"github.com/kataras/iris/v12"
)

func Service(app *iris.Application)  {
	app.PartyFunc("/wechat", Services.WechatService)
	app.PartyFunc("/web", Services.WebService)
}
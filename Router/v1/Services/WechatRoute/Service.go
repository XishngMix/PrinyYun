package WechatRoute

import (
	"PrintYun/Router/v1/View/WechatView"
	"github.com/kataras/iris/v12"
)

func User(party iris.Party) {
	party.Post("/login", WechatView.WechatLogin)
}
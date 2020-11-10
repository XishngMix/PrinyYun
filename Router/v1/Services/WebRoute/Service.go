package WebRoute

import (
	"PrintYun/Router/v1/View/WebView"
	"github.com/kataras/iris/v12"
)

func User(party iris.Party) {
	party.Post("/login", WebView.Login)
	party.Post("/register", WebView.Register)

}
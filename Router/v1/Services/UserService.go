package Services

import (
	"PrintYun/Router/v1/Services/WebRoute"
	"PrintYun/Router/v1/Services/WechatRoute"
	"github.com/kataras/iris/v12"
)

func WechatService(WechatParty iris.Party)  {
	WechatParty.PartyFunc("/user", WechatRoute.User)
}

func WebService(WebParty iris.Party)  {
	WebParty.PartyFunc("/user", WebRoute.User)
}

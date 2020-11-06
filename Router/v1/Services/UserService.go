package Services

import (
	"github.com/kataras/iris/v12"
	"PrintYun/Router/v1/View"
)

func UserService(UserParty iris.Party)  {
	UserParty.Post("/sendwxcode", View.WechatLogin)
}

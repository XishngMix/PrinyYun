package View

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"PrintYun/Router/v1/Structs"
	"PrintYun/ExternalCallCode"
)

func WechatLogin(ctx iris.Context) {
	WechatCode := Structs.WechatCode{}
	err := ctx.ReadForm(&WechatCode)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(fmt.Sprintf("{'message':'Sorry. An error occurred in the database. %s', Code='1001'}",  err.Error()))
	}
	_, err = ExternalCallCode.SendWxAuthAPI(WechatCode.Code)
	if err != nil {
		ctx.JSON("{'message':'Sorry. An error occurred in the database.', Code='1001'}")
	}
	ctx.JSON("{'message':'OK', Code='1000'}")
}
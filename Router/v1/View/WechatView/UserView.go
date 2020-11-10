package WechatView

import (
	"PrintYun/Router/v1/Structs"
	"PrintYun/middleware"
	"fmt"
	"github.com/kataras/iris/v12"
)

func WechatLogin(ctx iris.Context) {
	WechatCode := Structs.WechatCode{}
	err := ctx.ReadForm(&WechatCode)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(fmt.Sprintf("{'message':'抱歉,表单认证错误: %s', Code='1001'}",  err.Error()))
		return
	}
	//_, err = ExternalCallCode.SendWxAuthAPI(WechatCode.Code)
	//if err != nil {
	//	ctx.JSON("{'message':'Sorry. An error occurred in the database.', Code='1001'}")
	//	return
	//}
	tokenString := middleware.CreateJwt("1", "123")
	ctx.Header("Authorization", "bears "+tokenString)
	ctx.JSON("{'message':'OK', Code='1000'}")
}
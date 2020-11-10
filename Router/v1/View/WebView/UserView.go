package WebView

import (
	"PrintYun/DBSource"
	"PrintYun/DBSource/Modules"
	"PrintYun/ExternalCallCode"
	"PrintYun/Router/v1/Structs"
	"PrintYun/middleware"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"time"
)

func Login(ctx iris.Context)  {
	var form_Err, err error
	var db, data *gorm.DB
	LoginForm := Structs.ToLogin{}
	form_Err = ctx.ReadForm(&LoginForm)
	if form_Err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(fmt.Sprintf("{message:抱歉,表单认证错误: %s, Code='1001'}",  form_Err.Error()))
		return
	}

	defer ExternalCallCode.CollBackDB(db, err, ctx)
	db = DBSource.GetDB()
	data = db.Where(&Modules.Admin{UserName: LoginForm.UserName}).First(&Modules.Admin{})
	if data.RowsAffected == 0 {
		ctx.JSON(map[string]string{
			"message" : "未查询到此用户,请检验用户名后再次尝试",
			"code" : "1002",
		})
	}
	var maps Modules.Admin
	data.Scan(&maps)
	if ExternalCallCode.ComparePasswords(maps.PassWord, []byte(LoginForm.PassWord)) == false {
		ctx.JSON(map[string]string{
			"message" : "帐号或密码错误,请重新输入",
			"code" : "1005",
		})
	}
	token := middleware.CreateJwt2(maps.UserName, maps.Id)
	fmt.Println(token)

	ctx.JSON(map[string]string {
		"message":"登陆成功",
		"code" : "1000",
		"jwt": token,
	})
}

func Register(ctx iris.Context)  {
	var form_Err, err error
	var db *gorm.DB
	RegisterForm := Structs.Register{}
	form_Err = ctx.ReadForm(&RegisterForm)
	if form_Err != nil {
		ctx.JSON(map[string] string {
			"message": fmt.Sprintf("抱歉,表单认证错误: %s", form_Err.Error()),
			"code" : "1003",
		})
		return
	}

	// 查询数据库 对比验证码是否和数据库中的相同

	defer ExternalCallCode.CollBackDB(db, err, ctx)
	db = DBSource.GetDB()

	db.Create(&Modules.Admin{
		UserName: RegisterForm.UserName,
		PassWord: ExternalCallCode.HashAndSalt([]byte(RegisterForm.PassWord)),
		SignTime: time.Now().Unix(),
		SystemAuthority: 1,
	})
	//db.Commit()
	ctx.JSON(map[string]string {
		"message" : "操作成功",
		"code": "1000",
	})
}

//func VerificationCode(ctx iris.Context)  {
//
//}
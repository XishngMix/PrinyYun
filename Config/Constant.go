package Config


// 微信小程序系统相关
const (
	Code2SessiOnURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	APPID = "wx846a2a830abcf7b3"    						            // 小程序 appId
	SECRET = "851a94f375af4b306652d31089952269"							// 小程序 appSecretGRANT_TYPE = ""						// 授权类型，此处只需填写 authorization_code

	// 数据库设置相关
	MyUser = "xishng"
	Password = "Zxs781229"
	Host = "39.107.119.52"
	Port = 3306
	DbName = "PrintYun"

	JwtSECRET = "XishngPrintYun"
)
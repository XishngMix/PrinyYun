package UsuallyStruct

type ConfigReader struct {
	times int
}

//type JwtToken struct {
//	NickName string `json:"nick_name"`
//	//Emall string `json:"emall"`
//	Id int
//	//Iss string `json:"iss"`
//	//iat string `json:"iat"`
//	//jti string `json:"jti"`
//	//exp string `json:"exp"`
//	jwt.StandardClaims
//
//}
//    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//        "nick_name": "iris",
//        "email":"go-iris@qq.com",
//        "id":"1",
//        "iss":"Iris",
//        "iat":time.Now().Unix(),
//        "jti":"9527",
//        "exp":time.Now().Add(10*time.Hour * time.Duration(1)).Unix(),
//    })
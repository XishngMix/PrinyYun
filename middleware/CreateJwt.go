package middleware

import (
	"PrintYun/Config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/kataras/iris/v12"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func CreateJwt(Id, NIckname string) string {
	maxAge := 60*60*24
	claims := jwt.StandardClaims{
		Id : Id,
		Issuer: NIckname,
		ExpiresAt: time.Now().Add(time.Duration(maxAge)*time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Config.JwtSECRET))
	if err!=nil {
		fmt.Println(err)
	}
	ret,err := ParseToken(tokenString)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("userinfo: %v\n", ret)
	return tokenString
}


func ParseToken(tokenString string)(jwt.MapClaims,error)  {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(Config.JwtSECRET), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims,nil
	} else {
		return nil,err
	}
}

func JWT() iris.Handler {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte("XIShngPrintYun"), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)

		//debug 模式
		//Debug: bool
	})
	return jwtHandler.Serve
}

func CreateJwt2(NickName string, id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		       "nick_name": NickName,
		       "id": id,
		       "iss":"Iris",
		       "iat":time.Now().Unix(),
		       "jti":"9527",
		       "exp":time.Now().Add(10*time.Hour * time.Duration(1)).Unix(),
		   })
	tokenString, _ := token.SignedString([]byte("XIShngPrintYun"))
	return tokenString
}
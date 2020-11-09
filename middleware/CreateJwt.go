package middleware

import (
	"PrintYun/Config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
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
	fmt.Printf("token: %v\n", tokenString)
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
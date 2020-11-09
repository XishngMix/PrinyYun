package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func CrsAuth() iris.Handler {
	crs := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://foo.com"},   //允许通过的主机名称
		AllowCredentials: true,
	})

	return crs
}
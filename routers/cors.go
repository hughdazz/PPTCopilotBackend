package routers

import (
	"github.com/beego/beego/v2/server/web/context"
)

func cors_access(context *context.Context) {
	// 本地9529+7777端口
	context.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:9529")
	context.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "http://localhost:7777")
	context.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	context.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, token")
	context.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")

	// OPTIONS请求直接返回
	if context.Request.Method == "OPTIONS" {
		context.ResponseWriter.WriteHeader(200)
		return
	}
	// // 如果访问static目录下文件，直接返回
	// if context.Request.RequestURI[:7] == "/static" {
	// 	return
	// }

	// if context.Request.RequestURI == "/user/login" || context.Request.RequestURI == "/initial" {
	// 	return
	// }
	// _, err := context.Request.Cookie("token")
	// if err != nil {
	// 	context.ResponseWriter.WriteHeader(401)
	// 	return
	// }
}

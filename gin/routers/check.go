// Package routers @Author:冯铁城 [17615007230@163.com] 2023-06-26 15:29:30
package routers

import "go-study-net/gin/handler"

var checkParams1 = v1.Group("check-params")
var checkParams2 = v2.Group("check-params")

func initCheckRouter() {
	checkParams1.Use(handler.Interceptor())
	{
		checkParams1.GET("path-params-check/:id", handler.CheckPathParam)
		checkParams1.GET("url-params-check", handler.CheckUrlParam)
		checkParams1.POST("body-params-check", handler.CheckBodyJsonParam)
	}
	{
		checkParams2.GET("path-params-check/:id", handler.CheckPathParam)
		checkParams2.GET("url-params-check", handler.CheckUrlParam)
		checkParams2.POST("body-params-check", handler.CheckBodyJsonParam)
	}
}

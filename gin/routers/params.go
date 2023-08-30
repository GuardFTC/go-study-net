// Package routers @Author:冯铁城 [17615007230@163.com] 2023-06-26 15:15:01
package routers

import "go-study-net/gin/handler"

var params1 = v1.Group("params")
var params2 = v2.Group("params")

func initParamsRouter() {
	{
		params1.GET("headers", handler.GetHeader)
		params1.GET("path-params/:id", handler.GetPathParam)
		params1.GET("url-params", handler.GetUrlParams)
		params1.POST("body-params", handler.GetBodyJsonParam)
	}
	{
		params2.GET("path-params/:id", handler.GetPathParam)
		params2.GET("url-params", handler.GetUrlParams)
		params2.POST("body-params", handler.GetBodyJsonParam)
	}
}

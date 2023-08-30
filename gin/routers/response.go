// Package routers @Author:冯铁城 [17615007230@163.com] 2023-06-26 15:28:06
package routers

import "go-study-net/gin/handler"

var response1 = v1.Group("response")
var response2 = v2.Group("response")

func initResponseRouter() {
	{
		response1.GET("ok-response", handler.OkResponse)
		response1.GET("error-response", handler.ErrorResponse)
		response1.GET("json-response", handler.JsonResponse)
		response1.GET("xml-response", handler.XmlResponse)
		response1.GET("redirect-response", handler.RedirectResponse)
	}
	{
		response2.GET("ok-response", handler.OkResponse)
		response2.GET("error-response", handler.ErrorResponse)
		response2.GET("json-response", handler.JsonResponse)
		response2.GET("xml-response", handler.XmlResponse)
		response2.GET("redirect-response", handler.RedirectResponse)
	}
}

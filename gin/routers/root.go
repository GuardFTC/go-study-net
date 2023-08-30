// Package routers @Author:冯铁城 [17615007230@163.com] 2023-06-26 15:21:06
package routers

import "github.com/gin-gonic/gin"

// Router1 根服务,第1版本路由
var Router1 = gin.Default()
var v1 = Router1.Group("api/v1")

// Router2 根服务,第2版本路由
var Router2 = gin.Default()
var v2 = Router2.Group("api/v2")

// InitRouter 初始化路由
func InitRouter() {

	//1.初始化参数路由
	initParamsRouter()

	//2.初始化响应路由
	initResponseRouter()

	//3.初始化校验路由
	initCheckRouter()
}

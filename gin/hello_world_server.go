// Package gin @Author:冯铁城 [17615007230@163.com] 2023-07-25 09:14:38
package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartHelloWorldServer() {

	//1.初始化路由
	router := gin.Default()

	//2.绑定handler
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "hello world!",
		})
	})

	//3.开启服务器
	router.Run(":8080")
}

// Package handler Package gin @Author:冯铁城 [17615007230@163.com] 2023-06-26 14:46:21
package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Interceptor 自定义拦截器
func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		//1.获取请求头
		token := c.GetHeader("token")
		if token == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}

		//2.执行请求
		c.Next()

		//3.打印请求状态
		status := c.Writer.Status()
		log.Printf("request status is %v\n", status)
	}
}

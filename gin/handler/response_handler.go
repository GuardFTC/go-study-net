// Package handler @Author:冯铁城 [17615007230@163.com] 2023-06-26 09:23:33
package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// OkResponse 成功响应
func OkResponse(c *gin.Context) {

	//1.获取url参数
	id := c.Query("id")

	//2.响应给客户端
	c.String(http.StatusOK, "the id is %v\n", id)
}

// ErrorResponse 异常响应
func ErrorResponse(c *gin.Context) {

	//1.获取url参数
	code, _ := strconv.Atoi(c.Query("code"))

	//2.响应给客户端
	errorMessage := fmt.Sprintf("the error code is %v", code)
	c.AbortWithError(code, errors.New(errorMessage))
}

// JsonResponse json响应
func JsonResponse(c *gin.Context) {

	//1.拼接结构体切片
	students := []student{
		{Id: 1, Name: "ftc", Age: 18},
		{Id: 2, Name: "zyl", Age: 16},
	}

	//2.响应给客户端
	c.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}

// XmlResponse xml响应
func XmlResponse(c *gin.Context) {

	//1.拼接结构体
	s := &student{Id: 2, Name: "zyl", Age: 16}

	//2.响应给客户端
	c.XML(http.StatusOK, s)
}

// RedirectResponse 重定向响应
func RedirectResponse(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://www.baidu.com")
}

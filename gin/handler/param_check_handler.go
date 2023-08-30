// Package handler @Author:冯铁城 [17615007230@163.com] 2023-06-26 10:30:18
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 基础参数结构体
type baseParam struct {
	Id int `json:"id" uri:"id" binding:"required,min=1"`
}

// 分页基础参数结构体
type pageParam struct {
	PageNum  int `json:"pageNum" form:"pageNum" binding:"omitempty,min=1,max=10000"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"omitempty,min=1,max=100"`
}

// student查询参数结构体
type studentParam struct {
	Name string `json:"name" form:"name" binding:"omitempty,min=6,max=20"`
	Age  int    `json:"age" form:"age" binding:"omitempty,min=1,max=200"`
	pageParam
}

// CheckPathParam 校验路径参数
func CheckPathParam(c *gin.Context) {

	//1.声明结构体参数
	var param *baseParam

	//2.获取参数
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": param.Id,
		})
	}
}

// CheckUrlParam 校验URL参数
func CheckUrlParam(c *gin.Context) {

	//1.声明结构体参数
	var s *studentParam

	//2.获取参数
	if err := c.ShouldBindQuery(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": s,
		})
	}
}

// CheckBodyJsonParam 校验body参数
func CheckBodyJsonParam(c *gin.Context) {

	//1.声明结构体参数
	var s *studentParam

	//2.获取参数
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": s,
		})
	}
}

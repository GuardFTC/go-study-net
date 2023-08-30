// Package handler @Author:冯铁城 [17615007230@163.com] 2023-06-25 14:51:55
package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetHeader 获取请求头
func GetHeader(c *gin.Context) {

	//1.获取请求头
	token := c.GetHeader("token")

	//2.打印
	log.Printf("token is %v\n", token)
}

// GetPathParam 获取路径参数
func GetPathParam(c *gin.Context) {

	//1.获取参数
	id := c.Param("id")

	//2.打印
	log.Printf("path id is %v\n", id)
}

// student结构体变量
type student struct {
	Id   int    `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Age  int    `json:"age" form:"age" binding:"required"`
}

// GetUrlParams 获取URL参数
func GetUrlParams(c *gin.Context) {

	//1.获取非集合参数
	id := c.Query("id")
	log.Printf("url param id is %v\n", id)

	//2.获取非集合参数，当参数未空时，读取默认值
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	log.Printf("url param pageNum is %v,pageSize is %v\n", pageNum, pageSize)

	//3.获取集合参数
	ids := c.QueryArray("ids")
	for _, idsItem := range ids {
		log.Printf("url param ids is %v\n", idsItem)
	}

	//4.获取url参数-解析为结构体
	var s student
	if err := c.ShouldBindQuery(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("url param student is %v\n", s)
}

// GetBodyJsonParam 获取BODY参数
func GetBodyJsonParam(c *gin.Context) {

	//1.声明结构体变量
	var s student

	//2.将请求体数据绑定到结构体变量
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//3.打印json数据
	log.Printf("body param is %v\n", s)
}

// Package http_server @Author:冯铁城 [17615007230@163.com] 2023-06-21 15:15:00
package http_server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// 获取请求头handler
func getHeadHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取单个请求头
	token := req.Header.Get("token")
	log.Printf("request header token is [%v]\n", token)

	//2.获取所有请求头
	headers := req.Header
	for key := range headers {
		log.Printf("request header %v is %v\n", key, headers[key])
	}
}

// 获取路径参数handler
func getPathParamHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取请求路由
	route := req.URL.Path[1:]
	log.Printf("request routers is [%v]\n", route)

	//2.获取路由参数
	vars := mux.Vars(req)
	log.Printf("request path param is [%v]\n", vars["id"])
}

// 获取url参数handler
func getUrlParamHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取单个URL参数
	age := req.URL.Query().Get("age")
	log.Printf("request url param age is [%v]\n", age)

	//2.获取所有URL参数
	query := req.URL.Query()
	for key := range query {
		log.Printf("request url param %v is %v\n", key, query[key])
	}
}

// Student 定义student结构体
type Student struct {
	Id   int    `json:"id" validate:"required,gt=0,lte=10000"`
	Name string `json:"name" validate:"required,min=6,max=20"`
	Age  int    `json:"age" validate:"required,gt=0,lte=200"`
}

// 获取Body Json参数
func getBodyJsonParamHandler(w http.ResponseWriter, req *http.Request) {

	//1.读取body参数
	body, err := io.ReadAll(req.Body)
	if nil != err {
		log.Printf("get body error->%v\n", err.Error())
		return
	}

	//2.解析为JSON
	var student *Student
	if err = json.Unmarshal(body, &student); nil != err {
		log.Printf("parse json error->%v\n", err.Error())
		return
	}

	//3.打印参数
	log.Printf("request body param is %v\n", student)
}

// Package http_server @Author:冯铁城 [17615007230@163.com] 2023-06-21 16:00:53
package http_server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// OK响应处理器
func okResponseHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取URL参数name
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "新用户"
	}

	//2.响应
	w.Header().Set("Content-Type", "charset=utf-8")
	fmt.Fprint(w, "你好啊！"+name)
}

// 异常响应处理器
func errorResponseHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取URL参数name
	code, err := strconv.Atoi(req.URL.Query().Get("code"))
	if err != nil {
		log.Printf("get code error->[%v]\n", err.Error())
	}

	//2.响应
	switch code {
	case http.StatusNotFound:
		http.NotFound(w, req)
	default:
		http.Error(w, http.StatusText(code), code)
	}
}

// JSON响应处理器
func jsonResponseHandler(w http.ResponseWriter, req *http.Request) {

	//1.初始化结构体切片
	students := []*Student{
		{Id: 1, Name: "赵铁柱", Age: 98},
		{Id: 2, Name: "马冬梅", Age: 18},
	}

	//2.序列化
	result, err := json.Marshal(students)
	if err != nil {
		log.Printf("parse result error->[%v]\n", err.Error())
	}

	//3.返回
	w.Header().Set("Content-Type", "application/result;charset=utf-8")
	w.Write(result)
}

// 重定向处理器
func redirectHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://www.baidu.com/", http.StatusFound)
}

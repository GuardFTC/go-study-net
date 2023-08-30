// Package http_server @Author:冯铁城 [17615007230@163.com] 2023-06-25 14:08:07
package http_server

import (
	"net/http"
)

// 拦截器
func interceptor(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//1.获取请求头TOKEN参数
		token := req.Header.Get("token")
		if token == "" {
			http.Error(w, "token is nil", http.StatusUnauthorized)
			return
		}

		//2.执行请求
		next.ServeHTTP(w, req)

		//TODO 3.执行请求的后置处理
	}
}

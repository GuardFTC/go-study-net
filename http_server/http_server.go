// Package http_server @Author:冯铁城 [17615007230@163.com] 2023-06-19 17:22:41
package http_server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartHttpServer 开启服务器
func StartHttpServer() {

	route := mux.NewRouter()

	//1.加载获取参数相关Handler
	route.HandleFunc("/headers", getHeadHandler).Methods(http.MethodHead)
	route.HandleFunc("/path-params/{id}", getPathParamHandler).Methods(http.MethodGet)
	route.HandleFunc("/url-params", getUrlParamHandler).Methods(http.MethodGet)
	route.HandleFunc("/body-params", getBodyJsonParamHandler).Methods(http.MethodPost)

	//2.加载响应相关处理器
	route.HandleFunc("/ok-response", okResponseHandler).Methods(http.MethodGet)
	route.HandleFunc("/error-response", errorResponseHandler).Methods(http.MethodGet)
	route.HandleFunc("/json-response", jsonResponseHandler).Methods(http.MethodGet)
	route.HandleFunc("/redirect-response", redirectHandler).Methods(http.MethodGet)

	//3.加载参数校验处理器
	route.HandleFunc("/check-path-param/{id}", checkPathParamHandler).Methods(http.MethodGet)
	route.HandleFunc("/check-url-params", checkUrlParamHandler).Methods(http.MethodGet)
	route.HandleFunc("/check-body-params", checkBodyJsonParamHandler).Methods(http.MethodPost)

	//4.加载拦截器处理
	route.HandleFunc("/interceptor", interceptor(jsonResponseHandler)).Methods(http.MethodGet)

	//5.开启服务器并监听,如果出现异常,进行异常打印
	if err := http.ListenAndServe("127.0.0.1:8080", route); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

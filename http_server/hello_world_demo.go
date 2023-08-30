// Package http_server @Author:冯铁城 [17615007230@163.com] 2023-06-27 17:08:48
package http_server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StartHelloWorldServer hello world服务器
func StartHelloWorldServer() {

	//1.加载初始路由
	router := mux.NewRouter()

	//2.创建handler
	handler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "hello world!")
	}

	//3.绑定handler
	router.HandleFunc("/", handler).Methods(http.MethodGet)

	//4.开启服务器
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatalf("server run err->[%v]\n", err)
	}
}

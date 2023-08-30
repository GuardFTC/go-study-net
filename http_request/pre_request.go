// Package http_request @Author:冯铁城 [17615007230@163.com] 2023-06-27 10:59:23
package http_request

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// 初始化请求
func initRequest(method string, url string) *resty.Request {

	//1.创建客户端
	client := resty.New()

	//2.设置统一超时时间
	client.SetTimeout(10 * time.Second)

	//3.设置URL
	client.SetBaseURL(url)

	//4.初始化请求
	request := client.R()

	//5.设置请求方法
	request.Method = method

	//6.返回
	return request
}

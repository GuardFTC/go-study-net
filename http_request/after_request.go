// Package http_request @Author:冯铁城 [17615007230@163.com] 2023-06-27 14:26:16
package http_request

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// 处理响应结果
func dealResponse(resp *resty.Response) interface{} {

	//1.打印响应码以及响应体
	log.Printf("response code->[%v]\n", resp.StatusCode())
	log.Printf("response body->[%v]\n", string(resp.Body()))

	//2.校验响应码
	code := resp.StatusCode()
	if !(code >= http.StatusOK && code <= http.StatusIMUsed) {
		panic(fmt.Sprintf("request err code->[%v]\n", code))
	}

	//3.校验响应数据

	//4.返回
	return resp.Result()
}

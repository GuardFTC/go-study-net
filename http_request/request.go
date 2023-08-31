// Package http_request @Author:冯铁城 [17615007230@163.com] 2023-06-26 16:58:23
package http_request

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

var url = "https://jsonplaceholder.typicode.com/posts/1"
var client = resty.New()

// user结构体
type user struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

// GetRequest 发送GET请求
func GetRequest() {
	fmt.Println("-------------------------------------------------------------")

	//1.初始化请求
	request := initRequest(resty.MethodGet, url).
		SetQueryParams(map[string]string{
			"id":    "1",
			"name":  "ftc",
			"age":   "16",
			"title": "happy",
		}).
		SetHeader("Content-type", "application/json; charset=UTF-8").
		SetResult(&user{})

	//2.执行请求
	resp := doRequest(request)

	//3.处理请求结果
	dealResponse(resp)
}

// PostRequest 发送POST请求
func PostRequest() {
	fmt.Println("-------------------------------------------------------------")

	//1.初始化请求
	url := "https://jsonplaceholder.typicode.com/posts"
	request := initRequest(resty.MethodPost, url).
		SetBody(user{
			UserId: 2,
			Title:  "ftc",
			Body:   "ftc is the best",
		}).
		SetHeader("Content-type", "application/json; charset=UTF-8").
		SetResult(&user{})

	//2.执行请求
	resp := doRequest(request)

	//3.处理请求结果
	dealResponse(resp)
}

// PutRequest 发送PUT请求
func PutRequest() {
	fmt.Println("-------------------------------------------------------------")

	//1.初始化请求
	request := initRequest(resty.MethodPut, url).
		SetHeader("Content-type", "application/json; charset=UTF-8").
		SetBody(user{
			Id:     1,
			UserId: 1,
			Title:  "gn",
			Body:   "gn is the best",
		}).
		SetResult(&user{})

	//2.执行请求
	resp := doRequest(request)

	//3.处理请求结果
	dealResponse(resp)
}

// PatchRequest 发送PATCH请求
func PatchRequest() {
	fmt.Println("-------------------------------------------------------------")

	//1.初始化请求
	request := initRequest(resty.MethodPatch, url).
		SetHeader("Content-type", "application/json; charset=UTF-8").
		SetBody(user{
			Title: "wqw",
		}).
		SetResult(&user{})

	//2.执行请求
	resp := doRequest(request)

	//3.处理请求结果
	dealResponse(resp)
}

// DeleteRequest 发送DELETE请求
func DeleteRequest() {
	fmt.Println("-------------------------------------------------------------")

	//1.初始化请求
	request := initRequest(resty.MethodDelete, url)

	//2.执行请求
	resp := doRequest(request)

	//3.处理请求结果
	dealResponse(resp)
}

// UploadFileFromLocal 上传本地文件
func UploadFileFromLocal() {
	fmt.Println("-------------------------------------------------------------")

	//1.上传本地文件
	url := "https://myapp.com/upload"
	filePath := "/Users/a123/Documents/测试.png"
	request := initRequest(resty.MethodPost, url).
		SetFile("profile_img", filePath)

	//2.执行请求
	resp := doRequest(request)

	//3.处理请求结果
	dealResponse(resp)
}

// 执行请求
func doRequest(r *resty.Request) *resty.Response {

	//1.执行请求
	resp, err := r.Send()

	//2.打印入参
	log.Printf("request method->[%v]\n", r.Method)
	log.Printf("request url->[%v]\n", r.URL)
	log.Printf("request body->[%v]\n", r.Body)

	//3.校验异常
	if err != nil {
		panic(fmt.Sprintf("request err->[%v]\n", err.Error()))
	}

	//4.返回
	return resp
}
